module github.com/tortuoise/trac/cmd-gateway

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.8.3
	github.com/lib/pq v1.0.0
	github.com/tortuoise/trac/data v0.0.0
	github.com/tortuoise/trac/gateway v0.0.0
	github.com/tortuoise/trac/pb v0.0.0
	//golang.org/x/crypto v0.0.0-20190219172222-a4c6cb3142f2
	golang.org/x/net v0.0.0-20181220203305-927f97764cc3
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.19.0
)

replace github.com/tortuoise/trac/data => ../data

replace github.com/tortuoise/trac/pb => ../pb

replace github.com/tortuoise/trac/gateway => ../gateway

//replace github.com/grpc-ecosystem/grpc-gateway => $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway
