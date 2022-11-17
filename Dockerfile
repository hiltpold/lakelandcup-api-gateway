# Stage 1: compile the program
FROM golang:latest as build
WORKDIR /app
COPY go.* /app/
RUN go mod download
COPY . .
RUN go build -o server main.go

# Stage 2: build the image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app/
COPY --from=build /app/server .
COPY --from=build /app/config/envs/app.env /app/config/envs/
EXPOSE 50000
CMD ["./server"]  