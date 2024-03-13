BIN := bin
APP := foption
SRC := cmd/main.go

all: build

build: $(SRC)
	go build -o $(BIN)/$(APP) $(SRC)
