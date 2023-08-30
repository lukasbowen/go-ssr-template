build:
	@go build -o bin/server
run: build
	@./bin/server
styling-watch:
	@npx tailwindcss -i static/css/input.css -o static/css/style.css --watch
styling-min:
	@npx tailwindcss -i static/css/input.css -o static/css/style.css --minify
