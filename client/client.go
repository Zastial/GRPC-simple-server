package main

import (
	"context"
	"demo/invoicer"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()

	client := invoicer.NewInvoicerClient(conn)

	req := &invoicer.CreateRequest{
		Account: &invoicer.Account{
			Id:   2,
			Name: "Alexandre",
		},
		Request: &invoicer.RequestDest{
			From: &invoicer.Account{
				Id:   2,
				Name: "Alexandre",
			},
			To: &invoicer.Account{
				Id:   3,
				Name: "Antonin",
			},
		},
		Message: &invoicer.Message{
			Message: "Hello",
		},
	}
	response, err := client.GetInformations(context.Background(), req)
	if err != nil {
		log.Fatalf("could not create invoice: %v", err)
	}
	log.Printf("Response: %v", response)

}
