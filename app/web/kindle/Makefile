build:
	go build -o kindle cmd/main.go
	GOOS=linux GOARCH=amd64  go build
	docker build -t kindle .
