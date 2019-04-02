build:
	go build -o kindle main.go
	GOOS=linux GOARCH=amd64  go build
	docker build -t kindle .

run:
	docker run -it -p 8080:8080 kindle
