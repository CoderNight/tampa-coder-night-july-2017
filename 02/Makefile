.PHONY: default
default: run

missing-numbers: $(wildcard *.go) $(wildcard **/*.go)
	docker run --rm -v $(shell pwd):/src -w /src \
	-e GOOS=$(shell uname -s | tr '[:upper:]' '[:lower:]') -e GOARCH=amd64 \
	golang:alpine \
	go build -o missing-numbers

.PHONY: build
build: missing-numbers

.PHONY: run
run: missing-numbers
	@./missing-numbers

.PHONY: clean
clean:
	if [ -a missing-numbers ]; then rm missing-numbers; fi;

.PHONY: image
image: $(wildcard *.go) $(wildcard **/*.go)
	docker build -t joesteele/coder-night-missing-numbers .

.PHONY: run-image
run-image:
	@cat sample.txt | docker run -i --rm joesteele/coder-night-missing-numbers

.PHONY: clean-image
clean-image:
	docker rmi joesteele/coder-night-missing-numbers
