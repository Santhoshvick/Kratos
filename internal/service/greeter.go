package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	accountcustomerpb "account-customer/api/helloworld/v1"
	pb "account-service/api/helloworld/v1"
	"account-service/internal/biz"
	"account-service/internal/handler"
	customerpb "customer-service/api/helloworld/v1"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GreeterService is a greeter service.
type AccountService struct {
	pb.UnimplementedAccountServer
	customerClient customerpb.CustomerClient
	accountcustomerClient accountcustomerpb.AccountCustomerClient
	uc             *biz.AccountUsecase
}

// NewGreeterService new a greeter service.
func NewAccountService(uc *biz.AccountUsecase, custClient customerpb.CustomerClient,acccustClient accountcustomerpb.AccountCustomerClient) *AccountService {
	return &AccountService{uc: uc,
		customerClient: custClient,
		accountcustomerClient: acccustClient,

		
	}
}


func generateRefNumber() int {
	rand.Seed(time.Now().UnixNano()) 
	refNum := rand.Intn(9000000000) + 1000000000 
	return refNum
}
func connectToNATS() (*nats.Conn, error) {
	fmt.Println("Nats is connected successfully")
	return nats.Connect(nats.DefaultURL)
}

func (s *AccountService) CreateAccount(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {

	// Account Number Validation
	// AccountNumberValidation := strconv.FormatInt(req.AccountNumber, 10)
	// if len(AccountNumberValidation) < 10 || len(AccountNumberValidation) > 10 {
	// 	return &pb.CreateResponse{
	// 		Message: "The Digits of the Account number is not Valid Please provide 10-digit number ",
	// 	}, nil
	// }

	

	//Available Balance Validation
	if err:=handler.BalanceValidation(req.AvailableBalance); err!=nil{
		return &pb.CreateResponse{
			Message: ""+err.Error(),
		},nil
	}
	
	//Currency Validation 
	if err:=handler.CurrencyValidation(req.Currency); err!=nil{
		return &pb.CreateResponse{
			Message: ""+err.Error(),
		},nil
	}

	//Status Validation
	if err:=handler.StatusValidations(req.Status); err!=nil{
		return &pb.CreateResponse{
			Message: ""+err.Error(),
		},nil
	}

	g, err := s.uc.CreateAccount(ctx, &biz.Account{
		CustomerId:       req.CustomerId,
		AccountType:      req.AccountType,
		Currency:         req.Currency,
		Status:           req.Status,
		AvailableBalance: req.AvailableBalance,
		PendingBalance:   req.PendingBalance,
		CreditLimit:      req.CreditLimit,
	})

	if err != nil {
		return nil, err
	}
    
	//Customer Service GRPC Connection
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error Connecting: %v", err)
		return nil, err
	}
	client := customerpb.NewCustomerClient(conn)
	id, err := strconv.ParseInt(req.CustomerId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid customer ID: %w", err)
	}
	fmt.Println(id)

	val2, err := client.FindCustomer(ctx, &customerpb.FindRequest{CustomerId: int64(id)})
	if err != nil {
		log.Printf("Failed to find customer: %v", err)
		return nil, fmt.Errorf("customer with ID %d not found", id)
	}
	if val2 == nil || val2.CustomerId != id {
		log.Printf("Customer ID %d not found", id)
		return nil, fmt.Errorf("customer with ID %d does not exist", id)
	}
	nc, _ := connectToNATS()
	nc.Subscribe("Create", func(msg *nats.Msg) {
		fmt.Println("the account for the customer is created successfully", string(msg.Data))
	})

	randAccountNumber:=generateRefNumber()

	return &pb.CreateResponse{
		AccountId:        g.AccountId,
		CustomerId:       g.CustomerId,
		AccountNumber:    int64(randAccountNumber),
		AccountType:      g.AccountType,
		Currency:         g.Currency,
		Status:           g.Status,
		AvailableBalance: g.AvailableBalance,
		PendingBalance:   g.PendingBalance,
		CreditLimit:      g.CreditLimit,
		Message:          "Account For the Customer " + g.CustomerId + " is Created Successfully",
	}, nil
}



func (s *AccountService) UpdateAccount(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	g, err := s.uc.UpdateAccount(ctx, &biz.Account{
		AccountId:        req.AccountId,
		AccountNumber:    req.AccountNumber,
		AccountType:      req.AccountType,
		Currency:         req.Currency,
		Status:           req.Status,
		AvailableBalance: req.AvailableBalance,
		PendingBalance:   req.PendingBalance,
		CreditLimit:      req.CreditLimit,
	})
	if err != nil {
		return &pb.UpdateResponse{
			Message: "Couldn't Able to Update Because Customer Does not Exist",
		}, nil
	}
	return &pb.UpdateResponse{
		AccountId:        g.AccountId,
		CustomerId:       g.CustomerId,
		AccountNumber:    g.AccountNumber,
		AccountType:      g.AccountType,
		Currency:         g.Currency,
		Status:           g.Status,
		AvailableBalance: g.AvailableBalance,
		PendingBalance:   g.PendingBalance,
		CreditLimit:      g.CreditLimit,
	}, nil
}

func (s *AccountService) DeleteAccount(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	_, err := s.uc.DeleteAccount(ctx, req.AccountNumber)
	if err != nil {
		return &pb.DeleteResponse{
			Message: "Account Does Not Exist",
		}, nil
	}
	return &pb.DeleteResponse{
		Message: "Account is Deleted Successfully",
	}, nil
}

func (s *AccountService) FindAccount(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	g, err := s.uc.FindAccount(ctx, req.AccountId)
	if err != nil {
		return nil, err
	}
	return &pb.FindResponse{
		AccountId:        g.AccountId,
		CustomerId:       g.CustomerId,
		AccountNumber:    g.AccountNumber,
		AccountType:      g.AccountType,
		Currency:         g.Currency,
		Status:           g.Status,
		AvailableBalance: g.AvailableBalance,
		PendingBalance:   g.PendingBalance,
		CreditLimit:      g.CreditLimit,
	}, nil
}

func (s *AccountService) FindAccountNumber(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	g, err := s.uc.FindAccountNumber(ctx, req.AccountNumber)
	if err != nil {
		return &pb.FindResponse{
			Message: "Account Details Not Found",
		}, nil
	}
	return &pb.FindResponse{
		AccountId:        g.AccountId,
		CustomerId:       g.CustomerId,
		AccountNumber:    g.AccountNumber,
		AccountType:      g.AccountType,
		Currency:         g.Currency,
		Status:           g.Status,
		AvailableBalance: g.AvailableBalance,
		PendingBalance:   g.PendingBalance,
		CreditLimit:      g.CreditLimit,
		Message:          "Account Details Retrieved Successfully",
	}, nil
}

func (s *AccountService) FindByCustomer(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse1, error) {
	g, err := s.uc.FindByCustomer(ctx, req.CustomerId)
	if err != nil {
		return nil, err
	}
	fmt.Println(g)
	var pbAccounts []*pb.Account1
	for _, acc := range pbAccounts {
		pbAccounts = append(pbAccounts, &pb.Account1{
			AccountId:        acc.AccountId,
			CustomerId:       acc.CustomerId,
			AccountNumber:    acc.AccountNumber,
			AccountType:      acc.AccountType,
			Currency:         acc.Currency,
			Status:           acc.Status,
			AvailableBalance: acc.AvailableBalance,
			PendingBalance:   acc.PendingBalance,
			CreditLimit:      acc.CreditLimit,
		})
	}

	return &pb.FindResponse1{
		Account: pbAccounts,
	}, nil
}

func (s *AccountService) FindByCustomerId(ctx context.Context, req *pb.FindRequest1) (*pb.FindResponse2, error) {
	g, err := s.uc.FindByCustomerId(ctx, req.CustomerId)
	if err != nil {
		return &pb.FindResponse2{
			Message: "Account Does Not Exist",
		}, nil
	}
	return &pb.FindResponse2{
		AccountId:        g.AccountId,
		CustomerId:       g.CustomerId,
		AccountNumber:    g.AccountNumber,
		AccountType:      g.AccountType,
		Currency:         g.Currency,
		Status:           g.Status,
		AvailableBalance: g.AvailableBalance,
		PendingBalance:   g.PendingBalance,
		CreditLimit:      g.CreditLimit,
		Message:          "Account Exist",
	}, nil
}











// avalBalance, err := strconv.ParseInt(req.AvailableBalance, 10, 64)
	// if err != nil {
	// 	log.Fatalf("Issue with conveting the available Balance %v", err)
	// }

	// if avalBalance < 0 {
	// 	return &pb.CreateResponse{
	// 		Message: "Please Verify the Balance Once.Your Not allowed to make Your Balance negative",
	// 	}, nil
	// }

	// if req.Currency!="Rupee" && req.Currency!="USD" && req.Currency!="EURO" && req.Currency!="DIRHAM" {
	// 	return &pb.CreateResponse{
	// 		Message: "Please Verify the Currency Once.Allowed Curencies are USD  EURO  RUPEE   DIRHAM",
	// 	}, nil
	// }
	// if req.Status!="ACTIVE" && req.Status!="SUSPENDED" && req.Status!="CLOSED"{
	// 	return &pb.CreateResponse{
	// 		Message: "Please Verify the Status Of the account Once.Allowed Status are USD  ACTIVE   SUSPENDED   CLOSED",
	// 	}, nil
	// }



	//account-customer GRPC Connection
	// conn3, err := grpc.NewClient("localhost:9030", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Printf("Error Connecting: %v", err)
	// 	return nil, err
	// }
	// client2 := accountcustomerpb.NewAccountCustomerClient(conn3)

	// val3,err:=client2.CreateAccountCustomer(ctx,&accountcustomerpb.CreateCustomerAccountRequest{AccountId: req.AccountId,CustomerId: req.CustomerId})
	// if err!=nil{
	// 	return nil,err
	// }

	// fmt.Println(val3)

	// Account Number Validation
	// AccountNumberValidation := strconv.FormatInt(req.AccountNumber, 10)
	// if len(AccountNumberValidation) < 10 || len(AccountNumberValidation) > 10 {
	// 	return &pb.CreateResponse{
	// 		Message: "The Digits of the Account number is not Valid Please provide 10-digit number ",
	// 	}, nil
	// }