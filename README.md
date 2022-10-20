# horizontal-golang

一个用来表现 golang 在后台开发中效果的项目

## 环境

- golang 1.19
- gin 1.8.1

## 安装

```bash
$ go mod download
```

## 启动

```bash
# dev
$ go run main.go

# prod
$ export GIN_MODE=release
$ go build -o build/app
$ ./build/app
```
