package controller

import (
	"github.com/astaxie/beego/toolbox"
	pb "github.com/freeddser/pumpkin-rpc/proto/hello_http" // 引入proto包
	"golang.org/x/net/context"
	"time"
)

// defind helloHTTPService and implement the agreed interface.
type helloHTTPService struct{}

// HelloHTTPService Hello HTTP server
var HelloHTTPService = helloHTTPService{}

// SayHello implement the agreed interface.
func (h helloHTTPService) SayHelloPost(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	t := time.Now()
	resp := new(pb.HelloHTTPResponse)
	resp.Message = "Hello " + in.Name + "."

	elapsed := time.Duration(time.Since(t))
	toolbox.StatisticsMap.AddStatistics("POST", "/v1/echo", "&helloPBController.SayHelloPost", time.Duration(elapsed.Nanoseconds()))
	return resp, nil
}

// SayHello implement the agreed interface.
func (h helloHTTPService) SayHelloGet(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	t := time.Now()
	resp := new(pb.HelloHTTPResponse)
	resp.Message = "Hello " + in.Name + "."
	elapsed := time.Duration(time.Since(t))
	toolbox.StatisticsMap.AddStatistics("GET", "/v1/echo", "&helloPBController.SayHelloGet", time.Duration(elapsed.Nanoseconds()))
	return resp, nil
}
