# ace

一个使用了quic作为底层传输协议、包含http及rpc的后端框架，目的是建立高效高可靠的沙盒式后端框架。

## Quick start

### Requirments
Go version>=1.16

### Installation
```go

git clone https://github.com/Zenger-sun/ace.git

// 根据需要修改配置 
vim utils/cert/ca.conf
vim utils/cert/server.conf

sh utils/cert/gen.sh // 配置了默认值，执行时只需要一路回车

go mod download // 下载依赖包
go mod vendor // 将依赖复制到vendor下
```

### Build & Run

```go

go run server.go -I 127.0.0.1:8000
```

如出现秒退，请检查证书是否生成  