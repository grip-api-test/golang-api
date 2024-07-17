package service

import (
	"gin-gonic-api/app/domain/dao"
	"gin-gonic-api/app/repository"

	"golang.org/x/crypto/bcrypt"
)

type AccountService interface {
	GetAll() []dao.Account
	CreateAccount(newAccount dao.Account) dao.Account
	Get(id int) dao.Account
}

type LocalAccountServiceImpl struct {
	accountRepository repository.AccountRepository
}

func (u LocalAccountServiceImpl) CreateAccount(newAccount dao.Account) dao.Account {
	hash, _ := bcrypt.GenerateFromPassword([]byte(newAccount.Password), 15)
	newAccount.Password = string(hash)

	storedAccount := u.accountRepository.Save(&newAccount)

	return storedAccount
}

func (u LocalAccountServiceImpl) GetAll() []dao.Account {
	accounts := u.accountRepository.FindAll()
	return accounts
}

func AccountServiceInit(accountRepository repository.AccountRepository) *LocalAccountServiceImpl {
	return &LocalAccountServiceImpl{
		accountRepository: accountRepository,
	}
}

func (u LocalAccountServiceImpl) Get(id int) dao.Account {
	return u.accountRepository.Get(id)
}


