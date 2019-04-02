build:
	go build -o dogo main.go
	GOOS=linux GOARCH=amd64  go build
	docker build -t dogo .

run:
	docker run -it -p 8080:8080 blog
