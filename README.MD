# ProxyGet

> 自动从网上爬取一定量的Node并作测速

## 已有的ProxySource能获取到的协议

Sock5 MTProto SS SSR VMESS

## How To Add Source?

获取Source的方法在`service/internal/do/GetSub`中 对于使用已经存在的Clash订阅,往`Database/Database.db`的ClashSub下添加URL
对于是API的订阅,在`GetSub/API.go`中模仿已有的Func编写自己的Func,并在`Proxy.go`中添加

```go
Slice = append(Slice, YourFunc()...)
```

## Todo

- [x] 自动获取Proxy并且进行初步测试&存入数据库
- [ ] 生成ClashConfig
- [ ] 自动更新Proxy&[IP库](https://aite.xyz/share-file/qqwry/qqwry.dat)
- [ ] WebControlPanel

## Thank to

[RedElectricity](https://github.com/RedElectricity)
ProxyProviders
