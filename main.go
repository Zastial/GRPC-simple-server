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

/*
These are the functions that can be called by the client.
*/
func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	response := &invoicer.CreateResponse{
		Request: req.Request,
		Message: req.Message,
	}
	return response, nil
}

func (s myInvoicerServer) GetInformations(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.RequestDest, error) {

	response := &invoicer.RequestDest{
		From: req.Request.From,
		To:   req.Request.To,
	}
	return response, nil
}

/*
- main.go :
Launch a server on the port 8000.
Actually testing it with BloomRPC.

- client/client.go :
  - Launch a client on the port 8000.
  - Send a request to the server.
  - Print the response.
*/
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
