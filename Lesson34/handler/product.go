package handler

import (
    "encoding/json"
    "net/http"
    "strconv"

    "api/model"
    "api/storage/postgres"
    "github.com/go-chi/chi/v5"
)

type ProductHandler struct {
    Repo *postgres.ProductRepo
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product model.Product
    err := json.NewDecoder(r.Body).Decode(&product)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    err = h.Repo.CreateProduct(&product)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return 
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    product, err := h.Repo.GetProduct(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if product == nil {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    var product model.Product
    err = json.NewDecoder(r.Body).Decode(&product)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    product.ID = id

    err = h.Repo.UpdateProduct(&product)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    err = h.Repo.DeleteProduct(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
