## 1. client-go简介

​	client-go是一个调用kubernetes集群资源对象API的客户端，即通过client-go实现对kubernetes集群中资源对象（包括deployment、service、ingress、replicaSet、pod、namespace、node等）的增删改查等操作。大部分对kubernetes进行前置API封装的二次开发都通过client-go这个第三方包来实现。

​	client-go官方文档：https://github.com/kubernetes/client-go

## 2. client-go的使用

### 2.1 示例代码

```shell
git clone https://github.com/huweihuang/client-go.git
cd client-go
#保证本地HOME目录有配置kubernetes集群的配置文件
go run client-go.go
```

**[client-go.go](https://github.com/huweihuang/client-go/blob/master/client-go.go)**
