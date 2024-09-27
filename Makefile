build:
	@sudo go build  -o bin/main main.go
run: build
	@sudo  ./bin/main