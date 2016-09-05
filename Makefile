all: test

test:
	swatch

gofmt:
	gofmt -s -w .

lint: gofmt
