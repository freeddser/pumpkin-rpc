package controller

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego/toolbox"
	pb "github.com/freeddser/pumpkin-rpc/proto/monitor_http" // 引入proto包
	"golang.org/x/net/context"
	"truechain-engineering-code/log"
)

type monitorHTTPService struct{}

var MonitorHTTPService = monitorHTTPService{}

func (h monitorHTTPService) GetMonitors(ctx context.Context, in *pb.MonitorRequest) (*pb.MonitorResponse, error) {
	resp := new(pb.MonitorResponse)

	//monitor
	var goroutineInfo bytes.Buffer
	var heapInfo bytes.Buffer
	var threadcreateInfo bytes.Buffer
	var blockInfo bytes.Buffer
	var gcsummaryInfo bytes.Buffer
	toolbox.ProcessInput("lookup goroutine", &goroutineInfo)
	toolbox.ProcessInput("lookup heap", &heapInfo)
	toolbox.ProcessInput("lookup threadcreate", &threadcreateInfo)
	toolbox.ProcessInput("lookup block", &blockInfo)
	toolbox.ProcessInput("gc summary", &gcsummaryInfo)

	responseTime := toolbox.StatisticsMap.GetMapData()
	b, err := json.Marshal(responseTime)
	if err != nil {
		log.Error(err.Error())
	}

	//resturn infos
	resp.ResponseTime = string(b)
	resp.Goroutine = goroutineInfo.String()
	resp.Block = blockInfo.String()
	resp.Heap = heapInfo.String()
	resp.Threadcreate = threadcreateInfo.String()
	resp.GcSummary = gcsummaryInfo.String()

	return resp, nil
}
