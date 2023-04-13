package main

import (
	"context"
	"demo/invoicer"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	response := &invoicer.CreateResponse{
		Request: req.Request,
		Message: req.Message,
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("cannot create listener : %s", err)
	}

	serverRegister := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegister, service)
	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve : %s", err)
	}
}
