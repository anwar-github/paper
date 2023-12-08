FROM golang:1.20-alpine

# Set destination for COPY
WORKDIR /app

RUN apk add jq make git
RUN apk add --no-cache curl

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download  && go mod tidy && go mod verify

RUN go install github.com/githubnemo/CompileDaemon@latest

EXPOSE 8080

ENTRYPOINT CompileDaemon --build="go build -o build/paper" -command="./build/paper" -build-dir=/app