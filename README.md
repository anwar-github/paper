
# PAPER - TEST
The user has a balance in the application wallet and the balance wants to be
disbursed.
● Write code in Golang
● Write API (only 1 endpoint) for disbursement case only
● User data and balances can be stored as hard coded or database


## Installation

Install go https://go.dev/doc/install

```bash
  go version
  # go version go1.20.5 darwin/amd64
```

Install Docker https://docs.docker.com/engine/install

```bash
  docker -v
  # Docker version 24.0.2
  docker build -t referral:latest . 
```


## Run Locally

Clone the project

```bash
  git clone https://nobi-gitlab.usenobi.com/nobi-corp/referral.git
```

Go to the project directory

```bash
  cd referral
```

rename env-example

```bash
  cp env-example .env
```

Install dependencies

```bash
  go mod tidy
```

Start the server with build

```bash
  go run referral
```

Start the server with file main

```bash
  go run main.go
```


## Documentation

[ED Documentation](https://linktodocumentation)
[API Documentation](https://linktodocumentation)
## Tech Stack

**Client:** Swagger

**Server:** Go, Fiber, Gorm
## Authors

- [@platform](developer@usenobi.com)

