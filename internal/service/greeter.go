package service

import (
	"context"

	pb "account-service/api/helloworld/v1"
	"account-service/internal/biz"
)

// GreeterService is a greeter service.
type AccountService struct {
	pb.UnimplementedAccountServer

	uc *biz.AccountUsecase
}

// NewGreeterService new a greeter service.
func NewAccountService(uc *biz.AccountUsecase) *AccountService {
	return &AccountService{uc: uc}
}

func (s *AccountService) CreateAccount(ctx context.Context,req *pb.CreateRequest) (*pb.CreateResponse,error){
	g,err:=s.uc.CreateAccount(ctx,&biz.Account{
		AccountId:req.AccountId,
		AccountNumber:req.AccountNumber,
		AccountType:req.AccountType,
		Currency:req.Currency,
		Status:req.Status,
		AvailableBalance:req.AvailableBalance,
		PendingBalance:req.PendingBalance,
		CreditLimit:req.CreditLimit,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		AccountId: g.AccountId,
		AccountNumber: g.AccountNumber,
		AccountType:g.AccountType,
		Currency:g.Currency,
		Status:g.Status,
		AvailableBalance:g.AvailableBalance,
		PendingBalance:g.PendingBalance,
		CreditLimit:g.CreditLimit,
		
	},nil
}



func (s *AccountService) UpdateAccount(ctx context.Context,req *pb.UpdateRequest) (*pb.UpdateResponse,error){
	g,err:=s.uc.UpdateAccount(ctx,&biz.Account{
		AccountId:req.AccountId,
		AccountNumber:req.AccountNumber,
		AccountType:req.AccountType,
		Currency:req.Currency,
		Status:req.Status,
		AvailableBalance:req.AvailableBalance,
		PendingBalance:req.PendingBalance,
		CreditLimit:req.CreditLimit,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{
		AccountId: g.AccountId,
		AccountNumber: g.AccountNumber,
		AccountType:g.AccountType,
		Currency:g.Currency,
		Status:g.Status,
		AvailableBalance:g.AvailableBalance,
		PendingBalance:g.PendingBalance,
		CreditLimit:g.CreditLimit,
		
	},nil
}

func (s *AccountService) DeleteAccount(ctx context.Context,req *pb.DeleteRequest) (*pb.DeleteResponse,error){
	g,err:=s.uc.DeleteAccount(ctx,req.AccountId)
	if err != nil {
		return nil,err
	}

	return &pb.DeleteResponse{
		AccountId: g.AccountId,
		AccountNumber: g.AccountNumber,
		AccountType:g.AccountType,
		Currency:g.Currency,
		Status:g.Status,
		AvailableBalance:g.AvailableBalance,
		PendingBalance:g.PendingBalance,
		CreditLimit:g.CreditLimit,
		
	},nil
}

func (s *AccountService) FindAccount(ctx context.Context,req *pb.FindRequest) (*pb.FindResponse,error){
	g,err:=s.uc.FindAccount(ctx,req.AccountId)
	if err != nil {
		return nil, err
	}

	return &pb.FindResponse{
		AccountId: g.AccountId,
		AccountNumber: g.AccountNumber,
		AccountType:g.AccountType,
		Currency:g.Currency,
		Status:g.Status,
		AvailableBalance:g.AvailableBalance,
		PendingBalance:g.PendingBalance,
		CreditLimit:g.CreditLimit,
		
	},nil
}



