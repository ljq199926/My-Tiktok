# 抖音极简版

目录介绍：

apiGateway：网关

- deploy：放配置文件，yaml等
- handler：路由对应的处理的handler
- middleware：存放中间件，如跨域及jwt鉴权
- proto：存放每个微服务的密码本及生成的go文件，由微服务生成后copy过来即可
- router：路由初始化及路由匹配
- utils：存放工具类
- main.go：网关的入口函数
- go.mod：网关的包管理

Services：存放各个微服务，后期可拆分优化

- demoService
  - handler：服务的具体逻辑
  - proto：和网关存放的一致
  - subscriber：暂且未使用到
  - dockerfile：将微服务打包成镜像
  - makefile：需要安装MinGW才能使用make命令
  - main.go：微服务入口函数

demo运行：

先保证golang版本为1.14.1，如为高版本选择低版本进行安装即可，会自动卸载老版本，保证安装在同一目录，这样gopath等环境变量无需更改，高版本无法对micro v2进行支持，因为micro已经停止维护，拉取项目后，terminal使用命令

```
go get -v ./...
```

确保依赖都导入成功后，分别run两个入口文件

网关会在localhost:8080上运行

浏览器键入localhost:8080/demo/张三

网页上出现{"msg":"Hello 张三"}说明成功进行了http网关访问及网关与微服务的rpc调用

# 开发说明

首先cd到项目下边的Services，然后命令

```
micro new --gopath=false demoService(这个为微服务的名称)
```

然后micro工具集会自动生成这个微服务的基本框架，接着对demoService/proto/demoService下的proto文件进行修改，改为飞书文档里边的protobuf的数据传递格式，同时添加rpc方法，并加上一句option go_package = “proto/(demoService)”;

在有安装MinGW的情况下，直接执行

```
make proto
```

没有的情况下，打开Makefile文件，复制命令执行

```
protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/demoService/demoService.proto
```

#### 版本问题修改

打开生成的xxx.pb.micro.go文件，把import里面的包版本进行手动修改解决报错，加上v2

```go
client "github.com/micro/go-micro/v2/client"
server "github.com/micro/go-micro/v2/server"
```



接着修改handler文件下的demoService.go 文件，写微服务的具体逻辑实现

微服务端写完，把proto下的文件夹拷贝到网关的proto里面

然后在网关的router.go加入对应的路由，如注册的网关

```
route.GET("/douyin/user/register/",handler.Register)
```

然后去网关的handler里边添加对应的handler方法，比如上边的Register，写到userHandler.go里

后面路由多了就可以用路由组进行管理，然后在网关这边进行鉴权的操作

