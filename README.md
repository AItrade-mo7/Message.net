# Message.net

消息发送服务,用来发送消息，处理异步任务等。

## go.work

```go.work

go 1.20

use (
./
)

replace (
github.com/EasyGolang/goTools => /root/EasyGolang/goTools
)


```

## sys_env.yaml

```yaml
# 本地
MongoAddress: "127.0.0.1:27017"
MongoUserName: "mo7"
MongoPassword: "asdasd55555"
MessageBaseUrl: "http://127.0.0.1:8900"
# 运行模式
# RunMod: 1
```
