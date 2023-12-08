
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
  # Docker version 24.0.6
  docker-compose -v
  # Docker Compose version v2.21.0-desktop.1
  
```
Workspace Repository:
https://hub.docker.com/r/anwarmuhammad/paper/tags

## Run Locally

Clone the project

```bash
  git clone https://github.com/anwar-github/paper.git
```

Go to the project directory

```bash
  cd papper
```

rename env-example

```bash
  cp env-example .env
```

Start the server

```bash
  docker-compose up -d mysql
  # run app, port 8080
  docker-compose up -d app
```


## Api Definition
https://xkq7ys6pfk.apidog.io/

## Authors

- [Muhammad Anwar](myproject182@gmail.com)

