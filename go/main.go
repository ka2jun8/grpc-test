package main

import (
	"log"
	"net"

	pb "github.com/ka2jun8/grpc-test/go/proto"
	"github.com/ka2jun8/grpc-test/go/service"

	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	bookService := &service.BookService{}
	pb.RegisterBookServiceServer(server, bookService)
	server.Serve(listenPort)
}
