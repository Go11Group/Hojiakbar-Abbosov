package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var items = make(map[int]string)
var idCounter = 1

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getHandler(w, r)
	case "POST":
		postHandler(w, r)
	case "PUT":
		putHandler(w, r)
	case "DELETE":
		deleteHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	name, exists := items[id]
	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "ID: %d, Name: %s", id, name)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	id := idCounter
	idCounter++
	items[id] = name

	fmt.Fprintf(w, "Created Item - ID: %d, Name: %s", id, name)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	_, exists := items[id]
	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	items[id] = name

	fmt.Fprintf(w, "Updated Item - ID: %d, Name: %s", id, name)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, exists := items[id]
	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	delete(items, id)
	fmt.Fprintf(w, "Deleted Item - ID: %d", id)
}

func main() {
	http.HandleFunc("/items", handler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}


// GET /items?id=1 - Berilgan id bo'yicha elementni olish.
// POST /items?name=example - Yangi element yaratish.
// PUT /items?id=1&name=updated_example - Mavjud elementni yangilash.
// DELETE /items?id=1 - Mavjud elementni o'chirish.
// GET /items - Barcha elementlarni olish.
// POST /items - Body orqali yangi element yaratish.
// PUT /items - Body orqali mavjud elementni yangilash.
// DELETE /items - Barcha elementlarni o'chirish.
// GET /items/count - Elementlar sonini olish.
// GET /items/search?name=example - Nom bo'yicha elementni qidirish.