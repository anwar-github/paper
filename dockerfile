# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

# Set destination for COPY
WORKDIR /app

RUN apk add --no-cache curl

# Download Go modules
#COPY go.mod go.sum ./
RUN #go mod download  && go mod tidy && go mod verify

# Build wire depedency injector
# RUN go run github.com/google/wire/cmd/wire ./generator

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
#COPY *.go ./
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /papper

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/papper"]