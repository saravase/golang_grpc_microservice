gen:
	protoc --proto_path=proto/ proto/*.proto --go_out=plugins=grpc:pb
clean:
	rm pb/*.go
client:
	go run run/client/main.go -address 0.0.0.0:9090
server:
	go run run/server/main.go -port 9090
test:
	go test -cover -race ./...