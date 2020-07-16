module github.com/wmsx/store_svc

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.14
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/wmsx/xconf v0.0.0-20200710193800-f97c7e3c9e84
	google.golang.org/protobuf v1.25.0
)

replace (
	// 替换为v1.26.0版本的gRPC库
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
