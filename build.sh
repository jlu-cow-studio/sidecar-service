#!/bin/bash

# 替换为国内可用代理
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

# 构建应用
go build -o sidecar main.go