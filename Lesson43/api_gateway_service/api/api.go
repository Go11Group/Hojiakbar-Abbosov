package api

import (
	"net/http"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/api_gateway_service/api/handler"
	"github.com/gin-gonic/gin"
)

func Routes() *http.Server {
	mux := gin.Default()

	h := handler.NewHandler()

	mux.POST("/station/create", h.Client)

	return &http.Server{Handler: mux, Addr: ":8081"}
}
