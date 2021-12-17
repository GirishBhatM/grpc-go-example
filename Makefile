gen:
	protoc --proto_path=proto greeter.proto --go_out=. --go-grpc_out=.
clean:
	rm -rf pb/*
run-main:
	go run main.go
run-server:
	go run server/greetgrpcserver.go --port=50051
run-client:
	go run client/greeterclient.go --port=50051 --name=John