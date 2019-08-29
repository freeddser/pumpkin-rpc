package rpcserver

import (
	"github.com/freeddser/pumpkin-rpc/config"
	"github.com/freeddser/rs-common/util"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var GrpcServer *grpc.Server

var Endpoint string

//var Log = logging.MustGetLogger()
var SslCrtFile string
var SslKeyFile string
var SslDomain string

func init() {
	configFile := flag.String("c", "", "Configuration File")
	flag.Parse()

	if *configFile == "" {
		fmt.Println("\n\nUse -h to get more information on command line options\n")
		fmt.Println("You must specify a configuration file")
		os.Exit(1)
	}

	err := config.Initialize(*configFile)
	if err != nil {
		fmt.Printf("Error reading configuration: %s\n", err.Error())
		os.Exit(1)
	}

	SslCrtFile = config.MustGetString("pem.crt_file")
	SslKeyFile = config.MustGetString("pem.key_file")
	SslDomain = config.MustGetString("pem.domain")
	Endpoint = config.MustGetString("server.endpoint")

	util.InitTimeZoneLocation()

}

func GetGrpcServer() net.Listener {
	conn, err := net.Listen("tcp", Endpoint)
	if err != nil {
		grpclog.Fatalf("TCP Listen err:%v\n", err)
	}
	// grpc server
	creds, err := credentials.NewServerTLSFromFile(SslCrtFile, SslKeyFile)
	if err != nil {
		grpclog.Fatalf("Failed to create server TLS credentials %v", err)
	}
	GrpcServer = grpc.NewServer(grpc.Creds(creds))
	return conn
}

func GetTLSConfig() *tls.Config {
	cert, _ := ioutil.ReadFile(SslCrtFile)
	key, _ := ioutil.ReadFile(SslKeyFile)
	var demoKeyPair *tls.Certificate
	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		grpclog.Fatalf("TLS KeyPair err: %v\n", err)
	}
	demoKeyPair = &pair
	return &tls.Config{
		Certificates: []tls.Certificate{*demoKeyPair},
		NextProtos:   []string{http2.NextProtoTLS}, // HTTP2 TLS
	}
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func GrpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	if otherHandler == nil {
		return http.HandlerFunc(middlewareTime(func(w http.ResponseWriter, r *http.Request) {
			grpcServer.ServeHTTP(w, r)
		}))
	}
	return http.HandlerFunc(middlewareTime(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}))
}

func middlewareTime(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		defer func() {
			fmt.Printf("[%s] %s, time_used: %v", r.Method, r.URL.String(), time.Now().Sub(begin))
			log.Printf("[%s] %s, time_used: %v", r.Method, r.URL.String(), time.Now().Sub(begin))
		}()

		fn(w, r)
	}
}
