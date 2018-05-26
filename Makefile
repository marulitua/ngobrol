dev: clean serve

serve: build
	bin/chat

clean:
	rm -f bin/chat

build:
	go build -o bin/chat -v src/*.go
