package controller

import (
	"github.com/astaxie/beego/toolbox"
	pb "github.com/freeddser/pumpkin-rpc/proto/customer_http" // 引入proto包
	"github.com/freeddser/pumpkin-rpc/services"
	"golang.org/x/net/context"
	"time"
)

type customerHTTPService struct{}

var CustomerHTTPService = customerHTTPService{}

func (h customerHTTPService) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerGeneralResponse, error) {
	t := time.Now()
	success := true
	customerId := 1001
	customer, err := services.InsertCustomer(int64(customerId), in.Name, in.Email, in.Phone)
	if err != nil {
		success = false
		return nil, err
	}

	data := &pb.CustomerGeneralResponse{Success: success, CustomerInfos: []*pb.CustomerGeneralResponse_CustomerInfo{&pb.CustomerGeneralResponse_CustomerInfo{Id: customer.ID, Name: customer.Name, Email: customer.Email, Phone: customer.Phone}}}
	//monitor
	elapsed := time.Duration(time.Since(t))
	toolbox.StatisticsMap.AddStatistics("POST", "/v1/customer", "&customerPBController.CreateCustomer", time.Duration(elapsed.Nanoseconds()))

	return data, nil
}

func (h customerHTTPService) GetCustomers(ctx context.Context, in *pb.CustomerFilter) (*pb.CustomerList, error) {
	t := time.Now()
	customers := []*pb.CustomerList_Customer{}
	data, err := services.GetAllCustomers()
	if err != nil {
		return nil, err
	}

	for _, customer := range data {
		customers = append(customers, &pb.CustomerList_Customer{Id: customer.ID, Name: customer.Name, Email: customer.Email, Phone: customer.Phone})
	}

	//monitor
	elapsed := time.Duration(time.Since(t))
	toolbox.StatisticsMap.AddStatistics("GET", "/v1/customer/all", "&customerPBController.GetCustomers", time.Duration(elapsed.Nanoseconds()))

	return &pb.CustomerList{Customers: customers}, nil
}
