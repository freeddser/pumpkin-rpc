package controller

import (
	"github.com/astaxie/beego/toolbox"
	pb "github.com/freeddser/pumpkin-rpc/proto/echo_http" // 引入proto包
	"golang.org/x/net/context"
	"time"
)

// defind echoHTTPService and implement the agreed interface.
type echoHTTPService struct{}

// EchoHTTPService Echo HTTP server
var EchoHTTPService = echoHTTPService{}

// SayEcho implement the agreed interface.
func (h echoHTTPService) Echo(ctx context.Context, in *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	t := time.Now()

	resp := new(pb.SimpleMessage)
	resp.Id = "Echo " + in.Id + "."

	elapsed := time.Duration(time.Since(t))
	toolbox.StatisticsMap.AddStatistics("GET", "/v1/echo/{id}", "&echoPBController.Echo", time.Duration(elapsed.Nanoseconds()))

	return resp, nil
}

// SayEcho implement the agreed interface.
func (h echoHTTPService) EchoBody(ctx context.Context, in *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	t := time.Now()

	resp := new(pb.SimpleMessage)
	resp.Id = "Echo " + in.Id + "."

	elapsed := time.Duration(time.Since(t))
	toolbox.StatisticsMap.AddStatistics("POST", "/v1/echo_body", "&echoPBController.EchoBody", time.Duration(elapsed.Nanoseconds()))

	return resp, nil
}

func (h echoHTTPService) EchoDelete(ctx context.Context, in *pb.SimpleMessage) (*pb.SimpleMessage, error) {
	t := time.Now()

	resp := new(pb.SimpleMessage)
	resp.Id = "Echo " + in.Id + "."

	elapsed := time.Duration(time.Since(t))
	toolbox.StatisticsMap.AddStatistics("DELETE", "/v1/echo_delete/{id}", "&echoPBController.EchoDelete", time.Duration(elapsed.Nanoseconds()))

	return resp, nil
}
