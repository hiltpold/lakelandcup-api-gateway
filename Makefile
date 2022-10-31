proto:
	protoc services/**/pb/*.proto --go_out=. --go-grpc_out=.

api-gateway-dev:
	go run main.go