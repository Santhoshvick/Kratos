package biz

import (
	"context"
	"time"

	v1 "account-service/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Account struct {
	AccountId         int64 `gorm:"column:account_id;type:bigint;primaryKey"`
	CustomerId        string
	AccountNumber     int64    `gorm:"unique"`
	AccountType      string
	Currency          string
	AvailableBalance  string
	PendingBalance    string
	CreditLimit       string
	Status string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LastTransactionAt time.Time
}

// GreeterRepo is a Greater repo.
type AccountRepo interface {
	CreateAccount(context.Context, *Account) (*Account,error)
	UpdateAccount(context.Context, *Account) (*Account, error)
	DeleteAccount(context.Context, int64) (*Account, error)
	FindAccount(context.Context, int64) (*Account, error)
	FindAccountNumber(context.Context, int64) (*Account, error)
	FindByCustomer(context.Context, string) ([]*Account, error)
	FindByCustomerId(context.Context,string)(*Account,error)
}

// GreeterUsecase is a Greeter usecase.
type AccountUsecase struct {
	repo AccountRepo
	log  *log.Helper
}


func NewAccountUsecase(repo AccountRepo, logger log.Logger,) *AccountUsecase {
	return &AccountUsecase{repo: repo, log: log.NewHelper(logger)}
}


func (uc *AccountUsecase) CreateAccount(ctx context.Context, g *Account) (*Account,error) {
	return uc.repo.CreateAccount(ctx, g)
}

func (uc *AccountUsecase) UpdateAccount(ctx context.Context, g *Account) (*Account, error) {
	return uc.repo.UpdateAccount(ctx, g)
}

func (uc *AccountUsecase) DeleteAccount(ctx context.Context, accountId int64) (*Account, error) {
	return uc.repo.DeleteAccount(ctx, accountId)
}

func (uc *AccountUsecase) FindAccount(ctx context.Context,accountId int64) (*Account, error) {
	return uc.repo.FindAccount(ctx,accountId)
}


func (uc *AccountUsecase) FindAccountNumber(ctx context.Context,accountNumber int64) (*Account, error) {
	return uc.repo.FindAccountNumber(ctx,accountNumber)
}

func (uc *AccountUsecase) FindByCustomer(ctx context.Context,customerId string) ([]*Account, error) {
	return uc.repo.FindByCustomer(ctx,customerId)
}


func (uc *AccountUsecase) FindByCustomerId(ctx context.Context,customerId string) (*Account, error) {
	return uc.repo.FindByCustomerId(ctx,customerId)
}





