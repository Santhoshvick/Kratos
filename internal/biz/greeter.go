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
	AccountId         int64
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
	CreateAccount(context.Context, *Account) (*Account, error)
	UpdateAccount(context.Context, *Account) (*Account, error)
	DeleteAccount(context.Context, int64) (*Account, error)
	FindAccount(context.Context, int64) (*Account, error)
}

// GreeterUsecase is a Greeter usecase.
type AccountUsecase struct {
	repo AccountRepo
	log  *log.Helper
}


func NewAccountUsecase(repo AccountRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, log: log.NewHelper(logger)}
}


func (uc *AccountUsecase) CreateAccount(ctx context.Context, g *Account) (*Account, error) {
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



