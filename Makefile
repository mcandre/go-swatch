all: test

test:
	swatch

govet:
	go vet -v

gofmt:
	gofmt -s -w .

lint: govet gofmt
