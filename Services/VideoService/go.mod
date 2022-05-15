module VideoService

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo/redis v0.0.1
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/qiniu/go-sdk/v7 v7.11.1
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v2 v2.2.4
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
)
