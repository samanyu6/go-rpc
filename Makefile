gen: 
	protoc --proto_path=userpb .\userpb\user.proto --go_out=.

clean:
	rm grpc/*.go

run:
	go run main.go