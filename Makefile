# Windows
ifeq ($(OS),Windows_NT)
	EXTENSION=.exe
	RM=cmd /c rmdir /s /q
else
	EXTENSION=
	RM=rm -rf
endif

BIN=bin/swatch$(EXTENSION)

all: test

test: $(BIN)
	$(BIN)

$(BIN): cli.go
	go build -o $(BIN) cli.go

clean:
	-$(RM) bin
