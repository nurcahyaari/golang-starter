dev: generate
	go run github.com/cosmtrek/air

generate:
	go generate .

build:
	mkdir build && go build -o build/app