package service

import (
	"context"
	"errors"

	pb "github.com/ka2jun8/grpc-test/go/proto"
)

type BookService struct{}

func (s *BookService) GetBook(ctx context.Context, message *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	switch message.Id {
	case "book1":
		book := &pb.Book{
			Title:  "book1",
			Author: "author1",
		}
		return &pb.GetBookResponse{
			Book: book,
		}, nil
	}

	return nil, errors.New("Not Found book")
}
