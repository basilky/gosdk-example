.PHONY: all build env-up

all: build env-up run


##### BUILD
build:
	@echo "Build ..."
	@go build -o server
	@echo "Build done"

##### ENV
env-up:
	@echo "Start environment ..."
	@./startFabric.sh

##### RUN
run:
	@echo "Starting web app ..."
	@./server