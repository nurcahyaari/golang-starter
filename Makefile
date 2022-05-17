dev: generate
	run

run:
	go run github.com/cosmtrek/air

generate:
	go generate .

build: generate
	mkdir build && go build -o build/app