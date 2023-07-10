package main

import (
	"context"
	"log"
	"net"

	"github.com/saifuljnu/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoiceServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")

	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serviceRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}
	invoicer.RegisterInvoiceServer(serviceRegistrar, service)

	err = serviceRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve %s", err)
	}
}
