module github.com/liangzuobin/etcdttltest

go 1.14

require (
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/text v0.3.0
	google.golang.org/grpc v1.26.0
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
