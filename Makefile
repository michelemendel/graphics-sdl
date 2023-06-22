NAME=app

build:
	go build -o bin/$(NAME) -v

run: build
	./bin/$(NAME)

run_wo_build:
	go run main.go

test:
	go test -v ./...

clean:
	rm -rf bin