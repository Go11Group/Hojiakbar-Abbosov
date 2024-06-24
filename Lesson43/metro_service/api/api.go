package api

import (
	"database/sql"
	"net/http"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/metro_service/api/handler"
	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *http.Server {
	mux := gin.Default()

	h := handler.NewHandler(db)
	cardHandler := handler.NewCardHandler(db)
	transactionHandler := handler.NewTransactionHandler(db)
	terminalHandler := handler.NewTerminalHandler(db)

	mux.POST("/station/create", h.CreateStation)
	mux.GET("/station/:id", h.GetStationById)
	mux.PUT("/station/update", h.UpdateStation)
	mux.DELETE("/station/delete/:id", h.DeleteStation)

	mux.POST("/card/create", cardHandler.CreateCard)
	mux.GET("/card/:id", cardHandler.GetCardById)
	mux.PUT("/card/update", cardHandler.UpdateCard)
	mux.DELETE("/card/delete/:id", cardHandler.DeleteCard)

	mux.POST("/transaction/create", transactionHandler.CreateTransaction)
	mux.GET("/transaction/:id", transactionHandler.GetTransactionById)
	mux.PUT("/transaction/update", transactionHandler.UpdateTransaction)
	mux.DELETE("/transaction/delete/:id", transactionHandler.DeleteTransaction)

	mux.POST("/terminal/create", terminalHandler.CreateTerminal)
	mux.GET("/terminal/:id", terminalHandler.GetTerminalById)
	mux.PUT("/terminal/update", terminalHandler.UpdateTerminal)
	mux.DELETE("/terminal/delete/:id", terminalHandler.DeleteTerminal)

	return &http.Server{Handler: mux, Addr: ":8080"}
}
