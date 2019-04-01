.PHONY: all dev clean build env-up env-down run

all: clean build env-up run

dev: build run

##### BUILD
build:
	@echo "Build ..."
	@go build
	@echo "Build done"

##### ENV
env-up:
	@echo "Start environment ..."
	@cd fixtures && docker-compose -f docker-kafka.yaml up -d && docker-compose -f docker-zk.yaml up -d && docker-compose -f docker-compose-orderer.yaml up -d && docker-compose -f docker-compose-peer.yaml up -d
	@echo "Environment up"

env-down:
	@echo "Stop environment ..."
	@cd fixtures && docker-compose -f docker-kafka.yaml down && docker-compose -f docker-zk.yaml down && docker-compose -f docker-compose-orderer.yaml down && docker-compose -f docker-compose-peer.yaml down
	@echo "Environment down"

##### RUN
run:
	@echo "Start app ..."
	@./yunnan-project

##### CLEAN
clean: env-down
	@echo "Clean up ..."
	@rm -rf /tmp/kongyixueyuan-* kongyixueyuan
	@docker rm -f -v `docker ps -a --no-trunc | grep "kongyixueyuan" | cut -d ' ' -f 1` 2>/dev/null || true
	@docker rmi `docker images --no-trunc | grep "kongyixueyuan" | cut -d ' ' -f 1` 2>/dev/null || true
	@echo "Clean up done"
