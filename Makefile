# Makefile need to indent with tab instead of space
# indent with spaces lead to this error: Makefile:5: *** missing separator.  Stop.
SHELL := /bin/bash

# export all variable to sub Makefile as well
export

install:  ## install project dependencies
	go mod tidy
	go mod vendor

lint:
	golangci-lint run -v ./...

test:
	# run test for debezium-sink-server with 5 minutes timeout
	go test -timeout 300s -v ./...	

clean:  ## clean up build artifacts
	rm -rf ./dist ./build
	rm -f .*.out *.out
