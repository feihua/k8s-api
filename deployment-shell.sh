#!/bin/bash

#停止服务
docker stop k8s-test


#删除容器
docker rm k8s-test

#删除镜像
docker rmi k8s-test:v1

#删除none镜像
docker rmi $(docker images | grep "none" | awk '{print $3}')

#构建服务
docker build -t k8s-test:v1 -f Dockerfile .

#启动服务
docker run -itd --net=host --name=k8s-test k8s-test:v1
