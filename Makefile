.PHONY: all clean build env-up

all: clean env-down build env-up run

##### ENV DOWN
env-down:
	@echo "Removing docker containers"
	@cd network && echo y | ./byfn.sh down
#@cd network && echo y | ./byfn.sh rm

##### CLEAN
clean:
	@echo "Clean ..."
	@reset
	@rm -rf store

##### BUILD
build:
	@echo "Go build ..."
	@go build -o server
	@echo "Build done"

##### ENV UP
env-up:
	@echo "Start environment ..."
	@./startFabric.sh

##### RUN
run:
	@echo "Run  ..."
	@./server