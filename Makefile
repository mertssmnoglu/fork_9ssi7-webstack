.PHONY: build run test templ js-web-build js-web-watch js-admin-watch js-admin-build clean setup compose
MAKE = make

# Current project uses 'npm' as the package manager for the JS part.
# Options: npm, pnpm, yarn
NODE_PKG = npm

# Go commands
build:
	go build -o bin/app ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test ./...

templ:
	templ generate ./api/web/templates/

# JS commands
js-web-build:
	cd ./api/web && $(NODE_PKG) run build

js-web-watch:
	cd ./api/web && $(NODE_PKG) run watch

js-admin-watch:
	cd ./api/admin && $(NODE_PKG) run watch

js-admin-build:
	cd ./api/admin && $(NODE_PKG) run build

# Additional commands
clean:
	rm -rf bin/
	rm -rf tmp/
	rm -rf ./api/web/static/dist/
	rm -rf ./api/admin/dist/

setup:
	cd ./api/web && $(NODE_PKG) install
	cd ./api/admin && $(NODE_PKG) install
	go mod download
	$(MAKE) js-web-build
	$(MAKE) js-admin-build

compose:
	docker compose -f compose.dev.yml up --build
