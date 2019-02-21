build:
# 一定要注意 Makefile 中的缩进，否则 make build 可能报错 Nothing to be done for build
# # protoc 命令前边是一个 Tab，不是四个或八个空格
	protoc -I. --go_out=plugins=grpc:./ proto/email/email.proto
	
#编译docker镜像
	# GOOS=linux GOARCH=amd64 go build
	# go build -o order-service main.go
	# docker build -t order-service .

#运行docker镜像
# run:
# 	docker run -p 8001:8001 order-service

# stop:


# clean:
# 	rm order-service
# 	docker rmi order-service:latest