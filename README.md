# Calculator Service

This application provides a calculator service using a gRPC Unary API to perform basic arithmetic operations such as addition, subtraction, multiplication, and division on numbers received from users as inputs.

## Setup

This project requires a `gcc` compiler installed and the `protobuf` code generation tools.

### Install protobuf compiler

Install the `protoc` tool using the instructions available at [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/).

## Generating gRPC Code

To generate the gRPC code, run the following command:

```bash
protoc -Icalculator/proto --go_out=. --go_opt=module=github.com/AdekunleDally/grpc-unary-api --go-grpc_out=. --go-grpc_opt=module=github.com/AdekunleDally/grpc-unary-api calculator/proto/calculator.proto
