package main

import (
        "flag"
        "net/http"
        "os"

        "github.com/golang/glog"
        "golang.org/x/net/context"
        "github.com/grpc-ecosystem/grpc-gateway/runtime"
        "google.golang.org/grpc"

        pb "github.com/tortuoise/trac/pb"
        "github.com/tortuoise/trac/gateway"
)

var (
        tracEndpoint = flag.String("trac_endpoint", "localhost:9090", "endpoint of YourService")
	network    = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	swaggerDir = flag.String("swagger_dir", "/home/sridhar/dev/modgo/src/github.com/tortuoise/trac/pb", "path to the directory which contains swagger definitions")
)

func run() error {
        ctx := context.Background()
        ctx, cancel := context.WithCancel(ctx)
        defer cancel()

        mux := runtime.NewServeMux()
        opts := []grpc.DialOption{grpc.WithInsecure()}
        err := pb.RegisterTracHandlerFromEndpoint(ctx, mux, *tracEndpoint, opts)
        if err != nil {
                return err
        }

        return http.ListenAndServe(":8080", mux)
}

func main() {
        flag.Parse()
        defer glog.Flush()

        /*if err := run(); err != nil {
                glog.Fatal(err)
        }*/

        err := os.Chdir(*swaggerDir)
        if err != nil {
                glog.Errorf("%v", err)
        }

	ctx := context.Background()
	opts := gateway.Options{
		Addr: ":8080",
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *tracEndpoint,
		},
		SwaggerDir: *swaggerDir,
	}
	if err := gateway.Run(ctx, opts); err != nil {
		glog.Fatal(err)
	}
}
