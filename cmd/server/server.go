package main

import (
	"crypto/tls"
	"github.com/freeddser/pumpkin-rpc/config"
	"github.com/freeddser/pumpkin-rpc/repository"
	"github.com/freeddser/pumpkin-rpc/rpcserver"
	"github.com/freeddser/rs-common/logging"
	"github.com/freeddser/rs-common/util"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"net/http"
	"os"
	"time"
)

var log = logging.MustGetLogger()

func main() {
	start := util.GetTimeNow()
	//log
	setupLogging()

	//postgresql
	if config.MustGetString("switch.postgresql") == "on" {
		err := repository.InitFactory()
		if err != nil {
			log.Fatal("Cannot connect to database: ", err.Error())
			return
		}
	}

	//init conn rpc
	conn := rpcserver.GetGrpcServer()
	rpcserver.RegiserPBFromProto()

	// gateway server
	ctx := context.Background()
	dcreds, err := credentials.NewClientTLSFromFile(rpcserver.SslCrtFile, rpcserver.SslDomain)
	if err != nil {
		grpclog.Fatalf("Failed to create client TLS credentials %v", err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	gwmux := runtime.NewServeMux()
	rpcserver.RegiserPBFromEndpoint(ctx, gwmux, rpcserver.Endpoint, opts)

	// http server
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	srv := &http.Server{
		Addr:      rpcserver.Endpoint,
		Handler:   rpcserver.GrpcHandlerFunc(rpcserver.GrpcServer, mux),
		TLSConfig: rpcserver.GetTLSConfig(),
	}

	log.Info("Server Started in ", time.Since(start))
	log.Info("gRPC and https listen on: ", rpcserver.Endpoint)

	if err = srv.Serve(tls.NewListener(conn, srv.TLSConfig)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	return
}

func setupLogging() {
	logrus.SetLevel(logrus.DebugLevel)
	if config.MustGetString("switch.log") == "on" {
		log.Info("here")
		logPath := config.MustGetString("server.log_path")

		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
		if err != nil {
			log.Fatal("Cannot log to file", err.Error())
		}

		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(file)
	}
}
