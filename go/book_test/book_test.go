package test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mb "github.com/ka2jun8/grpc-test/go/mock_book"
	pb "github.com/ka2jun8/grpc-test/go/proto"
	"github.com/ka2jun8/grpc-test/go/service"
)

var book = &pb.Book{
	Title:  "book1",
	Author: "author1",
}

func TestGetBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBookClient := mb.NewMockBookServiceClient(ctrl)
	req := &pb.GetBookRequest{Id: "book1"}
	resp := &pb.GetBookResponse{Book: book}
	mockBookClient.EXPECT().GetBook(
		gomock.Any(),
		req,
	).Return(resp, nil)
	testGetBook(t, mockBookClient)
}

func testGetBook(t *testing.T, mock pb.BookServiceClient) {
	t.Helper()
	req := &pb.GetBookRequest{Id: "book1"}
	ctx := context.Background()
	expected, err := mock.GetBook(ctx, req)

	if err != nil {
		t.Errorf(err.Error())
	}
	bookService := &service.BookService{}
	actual, err2 := bookService.GetBook(ctx, req)
	if err2 != nil {
		t.Errorf(err2.Error())
	}

	if expected.Book.Title != actual.Book.Title {
		t.Errorf("mocking failed")
	}
	t.Log("Expect : ", expected)
	t.Log("Reply : ", actual)
}
