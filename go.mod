module dynamic-game

go 1.13

require (
	agones.dev/agones v1.4.0
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.3.5
	github.com/nats-io/nats.go v1.9.1
	github.com/pkg/errors v0.8.1
	k8s.io/api v0.0.0-20191004102255-dacd7df5a50b // kubernetes-1.12.10
	k8s.io/apiextensions-apiserver v0.0.0-20191004105443-a7d558db75c6 // kubernetes-1.12.10
	k8s.io/apimachinery v0.0.0-20191004074956-01f8b7d1121a // kubernetes-1.12.10
	k8s.io/client-go v9.0.0+incompatible // kubernetes-1.12.10
)
