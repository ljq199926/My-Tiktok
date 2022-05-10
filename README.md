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
  - dockerfile
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
