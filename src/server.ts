import * as grpc from "grpc";
import * as book_grpc_pb from "./proto/book_grpc_pb";
import BookService from "./book-service";

const server = new grpc.Server();
server.bind(
    `0.0.0.0:50051`,
    grpc.ServerCredentials.createInsecure(),
);

server.addService(
    book_grpc_pb.BookServiceService,
    new BookService(),
);

server.start();