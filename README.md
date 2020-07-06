# fabric-sdk-go-sample-gm
提供使用国密SDK和国密Fabric网络，国密Fabric-ca服务器进行交互的例子，提供通过API服务访问Fabric网络和Fabric-ca服务器。通过配置文件选择连接国密和非国密SDK。

# 开始

## 克隆项目源码
```
cd $GOPATH/src/github.com/hyperledger
git clone https://github.com/maluning/fabric-sdk-go-sample-gm.git
```

## 启动国密Fabric网络
```
cd $GOPATH/src/github.com/hyperledger/fabric/examples/network
./byfn.sh up -g

```
## 启动国密CA服务器
```
cd $GOPATH/src/github.com/hyperledger/fabric-ca/cmd/fabric-ca-server
go build
 ./fabric-ca-server start -b admin:adminpw

```
## 启动API服务
```
cd $GOPATH/src/github.com/hyperledger/fabric-sdk-go-sample-gm
go build
# 启动fabric服务
./fabric-sdk-go-sample-gm fabric
# 启动fabric-ca服务
./fabric-sdk-go-sample-gm ca

```
