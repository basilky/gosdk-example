.PHONY: all clean build env-up

all: clean env-down build env-up run

##### ENV DOWN
env-down:
	@cd network && echo y | ./byfn.sh down
#@cd network && echo y | ./byfn.sh rm

##### CLEAN
clean:
	@reset
	@echo "Clean ..."
	@rm -rf store

##### BUILD
build:
	@echo "Build ..."
	@go build -o server
	@echo "Build done"

##### ENV UP
env-up:
	@echo "Start environment ..."
	@./startFabric.sh

##### RUN
run:
	@echo "Starting web app ..."
	@./server