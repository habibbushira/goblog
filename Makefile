build:
	go build -o bin/goblog

run: build
	@./bin/goblog

watch:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go