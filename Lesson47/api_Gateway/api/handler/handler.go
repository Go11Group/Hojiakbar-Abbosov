package handler

import (
	"google.golang.org/grpc"
)

type handler struct {
	Connection grpc.ClientConn
}

func NewHandlerRepo(conn grpc.ClientConn) *handler {
	return &handler{Connection: conn}
}