package rpcserver

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pbHello "github.com/freeddser/pumpkin-rpc/proto/hello_http"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"

	"github.com/freeddser/pumpkin-rpc/controller"
	"github.com/freeddser/rs-common/logging"

	pbCustomer "github.com/freeddser/pumpkin-rpc/proto/customer_http"
	pbEcho "github.com/freeddser/pumpkin-rpc/proto/echo_http"
	pbMonitor "github.com/freeddser/pumpkin-rpc/proto/monitor_http"
)

var log = logging.MustGetLogger()

func RegiserPBFromProto() {
	pbHello.RegisterHelloHTTPServer(GrpcServer, controller.HelloHTTPService)
	pbEcho.RegisterEchoServiceServer(GrpcServer, controller.EchoHTTPService)
	pbCustomer.RegisterCustomerServiceServer(GrpcServer, controller.CustomerHTTPService)
	pbMonitor.RegisterMonitorServiceServer(GrpcServer, controller.MonitorHTTPService)
}

func RegiserPBFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	if err := pbHello.RegisterHelloHTTPHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		grpclog.Fatalf("Failed to register gw server: %v\n", err)
		log.Fatal("Failed to register gw server: %v\n", err)
	}

	if err := pbEcho.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		grpclog.Fatalf("Failed to register gw server: %v\n", err)
		log.Fatal("Failed to register gw server: %v\n", err)
	}

	if err := pbCustomer.RegisterCustomerServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		grpclog.Fatalf("Failed to register gw server: %v\n", err)
		log.Fatal("Failed to register gw server: %v\n", err)
	}

	if err := pbMonitor.RegisterMonitorServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		grpclog.Fatalf("Failed to register gw server: %v\n", err)
		log.Fatal("Failed to register gw server: %v\n", err)
	}

}
