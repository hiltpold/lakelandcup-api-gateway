proto:
	protoc services/**/pb/*.proto --go_out=. --go-grpc_out=.

dev-server:
	go run main.go