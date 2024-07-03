package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
    _ "github.com/lib/pq"
    "golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")
var db *sql.DB

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func initDB() {
    var err error
    connStr := "user=postgres password=root dbname=postgres sslmode=disable" 
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS refresh_tokens (
        id SERIAL PRIMARY KEY,
        username TEXT NOT NULL,
        token TEXT NOT NULL,
        expires_at TIMESTAMP NOT NULL,
        FOREIGN KEY (username) REFERENCES users (username)
    );`

    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    password := r.FormValue("password")

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    _, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
    if err != nil {
        http.Error(w, "User already exists", http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func VerifyJWT(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("Authorization")

        if tokenStr == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        claims := &Claims{}

        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    }
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    password := r.FormValue("password")

    var storedPassword string
    err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&storedPassword)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(15 * time.Minute)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenStr, err := token.SignedString(jwtKey)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    refreshToken := generateRefreshToken()
    refreshExpirationTime := time.Now().Add(24 * time.Hour)
    _, err = db.Exec("INSERT INTO refresh_tokens (username, token, expires_at) VALUES ($1, $2, $3)", username, refreshToken, refreshExpirationTime)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:    "access_token",
        Value:   tokenStr,
        Expires: expirationTime,
    })

    http.SetCookie(w, &http.Cookie{
        Name:    "refresh_token",
        Value:   refreshToken,
        Expires: refreshExpirationTime,
    })
}

func generateRefreshToken() string {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    refreshToken, _ := token.SignedString(jwtKey)
    return refreshToken
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("refresh_token")
    if err != nil {
        if err == http.ErrNoCookie {
            http.Error(w, "No refresh token provided", http.StatusUnauthorized)
            return
        }
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    refreshToken := cookie.Value

    var username string
    var expiresAt time.Time
    err = db.QueryRow("SELECT username, expires_at FROM refresh_tokens WHERE token = $1", refreshToken).Scan(&username, &expiresAt)
    if err != nil || time.Now().After(expiresAt) {
        http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(15 * time.Minute)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenStr, err := token.SignedString(jwtKey)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:    "access_token",
        Value:   tokenStr,
        Expires: expirationTime,
    })
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    initDB()
    defer db.Close()

    http.HandleFunc("/register", RegisterHandler)
    http.HandleFunc("/login", LoginHandler)
    http.HandleFunc("/refresh", RefreshHandler)
    http.HandleFunc("/hello", VerifyJWT(HelloHandler))

    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", nil)
}
