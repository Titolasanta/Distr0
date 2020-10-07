SHELL := /bin/bash
PWD := $(shell pwd)

GIT_REMOTE = github.com/7574-sistemas-distribuidos/docker-compose-init

default: build

all:

deps:
	go mod tidy
	go mod vendor

build: deps
	GOOS=linux go build -o bin/client github.com/7574-sistemas-distribuidos/docker-compose-init/client
.PHONY: build

docker-image:
	docker build -f ./server/Dockerfile -t "server:latest" .
	docker build -f ./client/Dockerfile -t "client:latest" .
.PHONY: docker-image

docker-compose-up: docker-image
	docker-compose -f docker-compose-dev.yaml up -d --build
.PHONY: docker-compose-up

docker-compose-up2: docker-image
	docker-compose -f docker-compose2.yaml up --scale client=2 -d --build
.PHONY: docker-compose-up

docker-compose-up: docker-image
	docker-compose -f docker-compose-dev.yaml up -d --build
.PHONY: docker-compose-up
docker-compose-down:
	docker-compose -f docker-compose-dev.yaml stop -t 1
	docker-compose -f docker-compose-dev.yaml down
.PHONY: docker-compose-down

docker-compose-vol-up: docker-image

	docker volume create config_vol
	docker build ./setup -t setup
	docker run --rm \
		--env CLI_SERVER_ADDRESS=server:11112 \
		--env SERVER_PORT=11112 \
		--env SERVER_LISTEN_BACKLOG=5 \
		--env CLI_ID=1 \
		--env CLI_LOOP_LAPSE=1m2s \
		--env CLI_LOOP_PERIOD=10s \
		-v config_vol:/data1 setup ./setup
	docker-compose -f docker-compose-dev2.yaml up -d --build
.PHONY: docker-compose-vol-up

docker-compose-vol-down:
	docker-compose -f docker-compose-dev2.yaml stop -t 1
	docker-compose -f docker-compose-dev2.yaml down
	docker volume rm config_vol
.PHONY: docker-compose-vol-down

docker-compose-down2:
	docker-compose -f docker-compose2.yaml stop -t 1
	docker-compose -f docker-compose2.yaml down
.PHONY: docker-compose-down2

docker-nc:

	docker-compose -f docker-compose-dev.yaml up -d --build
	docker network create  --subnet 172.25.126.0/24 red_netcat
	docker network connect --ip 172.25.126.2 red_netcat server
	docker build ./netcat/netcat/ -t netcat
	docker run --rm -it --network=red_netcat netcat echo "1" | nc -w 5 172.25.126.2 12345
.PHONY: docker-nc

docker-nc-clean:
	docker network disconnect red_netcat server
	docker network rm red_netcat
	docker-compose -f docker-compose-dev.yaml stop -t 1
	docker-compose -f docker-compose-dev.yaml down
.PHONY: docker-nc-clean

docker-compose-logs:
	docker-compose -f docker-compose-dev.yaml logs -f
.PHONY: docker-compose-logs

docker-compose-vol-logs:
	docker-compose -f docker-compose-dev2.yaml logs -f
.PHONY: docker-compose-vol-logs
