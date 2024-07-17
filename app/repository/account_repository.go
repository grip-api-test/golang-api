package repository

import (
	"gin-gonic-api/app/domain/dao"
)

type AccountRepository interface {
	FindAll() []dao.Account
	Save(account *dao.Account) dao.Account
}

type FakeAccountRepositoryImpl struct {
	accounts []dao.Account
}

func (u *FakeAccountRepositoryImpl) FindAll() []dao.Account {
	return u.accounts
}

func (u *FakeAccountRepositoryImpl) Save(account *dao.Account) dao.Account {
	u.accounts = append(u.accounts, *account)
	return *account
}

func AccountRepositoryInit() *FakeAccountRepositoryImpl {
	return &FakeAccountRepositoryImpl{
		accounts: []dao.Account{},
	}
}
