clean:
	rm sample/*.go

build:
	go build main.go

run: all
	go run main.go

gen:
	protoc --proto_path=sample --go_out=sample --go_opt=paths=source_relative example.proto

all: clean gen build

.PHONY: all