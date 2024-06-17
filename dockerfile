# Dockerfile

FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o http ./cmd/http/main.go
RUN go build -o grpc ./cmd/grpc/main.go
RUN go build -o graphql ./cmd/graph/main.go

EXPOSE 8080
EXPOSE 50051
EXPOSE 8081

CMD ["sh", "-c", "./http & ./grpc & ./graphql"]
