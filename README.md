# go-rpc

Super simple boilerplate project for grpc with Golang.

- Generating go Code from proto file 
```go
   protoc --proto_path=userpb .\userpb\user.proto --go_out=plugins=grpc:.
```