FROM golang:1.15 AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex
EXPOSE 8080
CMD reflex -g '*.go' go run cmd/web/main.go cmd/web/middleware.go cmd/web/routes.go --start-service

