package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/freeddser/pumpkin-rpc/config"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var Endpoint string

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

	Endpoint = config.MustGetString("server.endpoint")

}
func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://" + Endpoint + "/v1/monitor")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var rs AutoGenerated
	err = json.Unmarshal(body, &rs)
	if err != nil {
		fmt.Println(err)
	}

	var avgs []ResponseTime
	err = json.Unmarshal([]byte(rs.ResponseList), &avgs)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rs.Goroutine)
	fmt.Println(rs.Heap)
	fmt.Println(rs.Threadcreate)
	fmt.Println(rs.Block)
	fmt.Println(rs.GcSummary)
	fmt.Println("########### Response Time List #######################################")
	for _, avg := range avgs {
		fmt.Println("--------------------------------------------------------------------")
		fmt.Println("|avg_time:" + avg.AvgTime + "  |max_time:" + avg.MaxTime + "  |method:" + avg.Method + "  |min_time:" + avg.MinTime + "  |request_url:" + avg.RequestURL + "  |times:" + strconv.Itoa(avg.Times) + "  |total_time:" + avg.TotalTime + "|")

	}
}

type AutoGenerated struct {
	Heap         string `json:"heap"`
	Threadcreate string `json:"threadcreate"`
	Block        string `json:"block"`
	Goroutine    string `json:"goroutine"`
	GcSummary    string `json:"gc_summary"`
	ResponseList string `json:"response_time"`
}

type ResponseTime struct {
	AvgTime    string `json:"avg_time"`
	MaxTime    string `json:"max_time"`
	Method     string `json:"method"`
	MinTime    string `json:"min_time"`
	RequestURL string `json:"request_url"`
	Times      int    `json:"times"`
	TotalTime  string `json:"total_time"`
}
