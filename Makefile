GO = @go
NPM = @npm
NODEMON = @nodemon
DOCKER = @docker
COMPOSE = @docker-compose

#################################
# Application Territory
#################################
.PHONY: install
install:
	${GO} get .
	${GO} mod verify
	${NPM} i nodemon@latest -g

.PHONY: dev
dev:
	${NODEMON} -V -e .go,.env -w . -x go run ./cmd/api --count=1 --race -V --signal SIGTERM

.PHONY: build
build:
	${GO} mod tidy
	${GO} mod verify
	${GO} vet --race -v .
	${GO} build --race -v -o ${type} .

.PHONY: test
test:
	${GO} test -v ./domain/services/**

#################################
# Docker Territory
#################################
build:
	${DOCKER} build -t go-api:latest --compress .

up:
	${COMPOSE} up -d --remove-orphans --no-deps --build

down:
	${COMPOSE} down