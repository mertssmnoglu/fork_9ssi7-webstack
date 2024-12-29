.PHONY: build run dev clean js-build js-watch templ compose

# Go commands
build:
	go build -o bin/app ./cmd/main.go

run:
	go run ./cmd/main.go

js-build:
	cd app/api/web && npm run build

templ:
	templ generate ./api/web/templates/

js-watch:
	cd app/api/web && npm run watch

clean:
	rm -rf bin/
	rm -rf app/api/web/static/js/dist/

setup:
	cd app/api/web && npm install
	go mod download

test:
	go test ./...

compose:
	docker compose -f compose.dev.yml up --build