build:
	go build -o kindle cmd/main.go
	GOOS=linux GOARCH=amd64
	docker build -t kindle .
