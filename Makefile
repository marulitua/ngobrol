NODE_DIR := ./node_modules

dev: clean serve

serve: build build_assets
	bin/chat

clean:
	rm -f bin/chat

build:
	go build -o bin/chat -v src/backend/*.go

build_assets: | $(NODE_DIR)
	node node_modules/parcel/bin/cli.js build -d public src/frontend/index.html

serve_assets: | $(NODE_DIR)
	node node_modules/parcel/bin/cli.js serve -d public src/frontend/index.html

clean_assets:
	rm -rf public/*

$(NODE_DIR): ;@echo "Instaling node dependency using npm...";
	npm i
	go get -u "github.com/gorilla/websocket"
