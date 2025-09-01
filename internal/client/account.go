package client

import (

	customerpb "customer-service/api/helloworld/v1"
	accountcustomerpb "account-customer/api/helloworld/v1"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

var endpoint string = "0.0.0.0:9000"

func NewCustomerClient(endpoint string) (customerpb.CustomerClient, error) {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint))
	if err != nil {
		return nil, err
	}
	return customerpb.NewCustomerClient(conn), nil	
}


var endpoint1 string = "0.0.0.0:9030"

func NewAccountCustomerClient(endpoint1 string) (accountcustomerpb.AccountCustomerClient, error) {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint))
	if err != nil {
		return nil, err
	}
	return accountcustomerpb.NewAccountCustomerClient(conn), nil	
}

