.PHONY: default
default: build

.PHONY: build
build:
	go build -o bin/api ./go

.PHONY: test
test:
	go test github.com/ka2jun8/grpc-test/go/book_test -v

.PHONY: run/local
run/local:
	go run go/main.go

.PHONY: protoc
protoc:
	protoc --go_out=plugins=grpc:./go proto/book.proto

.PHONY: doc
doc:
	protoc --doc_out=html,index.html:./ proto/*.proto
