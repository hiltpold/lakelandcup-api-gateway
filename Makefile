proto:
	protoc services/**/pb/*.proto --go_out=. --go-grpc_out=.

build:
	go build -ldflags "-X github.com/hiltpold/lakelandcup-fantasy-service/commands.Version=`git rev-parse HEAD`"

api-gateway-dev:
	go run main.go -c .dev.env

api-gateway-prod:
	go run main.go -c .prod.env