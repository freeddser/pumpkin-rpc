package main

import (
	pb "github.com/freeddser/pumpkin-rpc/proto/hello_http" // 引入proto包

	pb2 "github.com/freeddser/pumpkin-rpc/proto/customer_http" //

	"flag"
	"fmt"
	"github.com/freeddser/pumpkin-rpc/config"
	"github.com/freeddser/rs-common/logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/grpclog"
	"os"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50055"
)

var log = logging.MustGetLogger()
var sslCrtFile string
var sslDomain string

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

	sslCrtFile = config.MustGetString("pem.crt_file")
	sslDomain = config.MustGetString("pem.domain")

}

func main() {
	// TLS连接
	creds, err := credentials.NewClientTLSFromFile(sslCrtFile, sslDomain)
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloHTTPClient(conn)

	// 调用方法
	reqBody := new(pb.HelloHTTPRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHelloGet(context.Background(), reqBody)
	if err != nil {
		grpclog.Fatalln(err)
	}

	log.Info(r.Message)

	//customer
	//pb2 "github.com/freeddser/pumpkin-rpc/proto/customer_http" ////customer

	customeReq := pb2.CustomerRequest{Id: 103, Name: "gavin1", Email: "gavin@aa.com", Phone: "1010101", Addresses: []*pb2.CustomerRequest_Address{
		&pb2.CustomerRequest_Address{
			Street:            "1 Mission Street",
			City:              "San Francisco",
			State:             "CA",
			Zip:               "94105",
			IsShippingAddress: false,
		},
		&pb2.CustomerRequest_Address{
			Street:            "Greenfield",
			City:              "Kochi",
			State:             "KL",
			Zip:               "68356",
			IsShippingAddress: true,
		}}}

	c1 := pb2.NewCustomerServiceClient(conn)
	r1, err := c1.CreateCustomer(context.Background(), &customeReq)
	fmt.Println(r1)

	fmt.Println("--------------")
	filter := &pb2.CustomerFilter{Keyword: "gavin"}
	r2, err := c1.GetCustomers(context.Background(), filter)
	fmt.Println(r2)
	fmt.Println("=========")

}

// OR: curl -X POST -k https://localhost:50052/example/echo -d '{"name": "gRPC-HTTP is working!"}'
//curl -X GET -k 'https://localhost:50052/example/echo?name=xx'
