package data

import (
	"context"

	"account-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type AccountRepo struct {
	data  *Data
	log   *log.Helper
	table *gorm.DB
}

// NewGreeterRepo .
func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &AccountRepo{
		data:  data,
		log:   log.NewHelper(logger),
		table: data.db.Table("account"),
	}
}

func (r *AccountRepo) CreateAccount(ctx context.Context, g *biz.Account) (*biz.Account, error) {
	result := r.table.WithContext(ctx).Create(g)
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}

func (r *AccountRepo) UpdateAccount(ctx context.Context, g *biz.Account) (*biz.Account, error) {
	result := r.table.Model(g).Where("account_number = ?", g.AccountNumber).Updates(biz.Account{AccountId: g.AccountId, AccountNumber: g.AccountNumber, Currency: g.Currency, CustomerId: g.CustomerId, AvailableBalance: g.AvailableBalance, PendingBalance: g.PendingBalance})
	if result.Error != nil {
		return nil, result.Error
	}
	return g, nil
}

func (r *AccountRepo) DeleteAccount(ctx context.Context, accountNumber int64) (*biz.Account, error) {
	result := r.table.WithContext(ctx).Where("account_number = ?", accountNumber).Delete(&biz.Account{})
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Account{}, nil
}

func (r *AccountRepo) FindAccount(ctx context.Context, accountId int64) (*biz.Account, error) {
	var acc biz.Account
	result := r.table.WithContext(ctx).Where("account_id = ?", accountId).First(&acc)
	if result.Error != nil {
		return nil, result.Error
	}

	return &acc, nil
}


func (r *AccountRepo) FindAccountNumber(ctx context.Context, accountNumber int64) (*biz.Account, error) {
	var acc biz.Account
	result := r.table.WithContext(ctx).Where("account_number = ?", accountNumber).First(&acc)
	if result.Error != nil {
		return nil, result.Error
	}

	return &acc, nil
}
func (r *AccountRepo) GetBalance(ctx context.Context, accountId int64) (*biz.Account, error) {
	var acc biz.Account
	result := r.table.WithContext(ctx).Where("account_id = ?", accountId).First(&acc)
	if result.Error != nil {
		return nil, result.Error
	}

	return &acc, nil
}


func (r *AccountRepo) FindByCustomer(ctx context.Context, customerId string) ([]*biz.Account, error) {
	var acc []*biz.Account
	result := r.table.WithContext(ctx).Where("customer_id = ?",customerId).Find(&acc)
	if result.Error != nil {
		return nil, result.Error
	}

	return acc, nil
}

func (r *AccountRepo) FindByCustomerId(ctx context.Context, customerId string) (*biz.Account, error) {
	var acc *biz.Account
	result := r.table.WithContext(ctx).Where("customer_id = ?",customerId).First(&acc)
	if result.Error != nil {
		return nil, result.Error
	}

	return acc, nil
}


