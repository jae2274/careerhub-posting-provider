BINARY_NAME=myapp
CODE_DIR=./tft/producer

include test.env

env:
	@env API_KEY=${API_KEY} DB_ENDPOINT=${DB_ENDPOINT} SQS_ENDPOINT=${SQS_ENDPOINT} QUEUE_NAME=${QUEUE_NAME}

## build: Build binary
build:
	@echo "Building..."
	@go build -ldflags="-s -w" -o ${BINARY_NAME} ${CODE_DIR}
	@echo "Built!"

## run: builds and runs the application
run: build	env
	@echo "Starting..."
	@./${BINARY_NAME} 
	@echo "Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

## test: runs all tests
test:	env
	@go test -p 1 -timeout 30s ./test/...

