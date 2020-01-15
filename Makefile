.PHONY: all clean build env-up

all: clean build env-up run

##### CLEAN
clean:
	@echo "Clean ..."
	@rm -rf store

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