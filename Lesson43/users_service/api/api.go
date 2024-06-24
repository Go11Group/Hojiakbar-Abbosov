package api

import (
	"database/sql"
	"net/http"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/users_service/api/handler"
	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *http.Server {
	mux := gin.Default()

	h := handler.NewHandler(db)
	mux.POST("/user/create", h.CreateUser)
	mux.GET("/user/:id", h.GetUserById)
	mux.PUT("/user/update", h.UpdateUser)
	mux.DELETE("/user/delete/:id", h.DeleteUser)

	return &http.Server{Handler: mux, Addr: ":8080"}
}
