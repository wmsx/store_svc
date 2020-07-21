module github.com/wmsx/store_svc

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.5.0
	github.com/wmsx/xconf v0.0.0-20200710193800-f97c7e3c9e84
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/mysql v0.3.0
	gorm.io/gorm v0.2.23
)

// 替换为v1.26.0版本的gRPC库
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
