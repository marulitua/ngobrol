NODE_DIR := ./node_modules

dev: clean serve

serve: build
	bin/chat

clean:
	rm -f bin/chat

build:
	go build -o bin/chat -v src/backend/*.go

build_assets:
	node node_modules/parcel/bin/cli.js build

serve_assets: | $(NODE_DIR)
	node node_modules/parcel/bin/cli.js src/frontend/index.html

clean_assets:
	rm -rf dist

$(NODE_DIR): ;@echo "Instaling node dependency using npm...";
	npm i
