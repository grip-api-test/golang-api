package controller

import (
	"gin-gonic-api/app/domain/dao"

	"github.com/stretchr/testify/mock"
)

type MockAccountService struct {
	mock.Mock
}

func (m *MockAccountService) GetAll() []dao.Account {
	args := m.Called()
	return args.Get(0).([]dao.Account)
}

func (m *MockAccountService) CreateAccount(newAccount dao.Account) dao.Account {
	args := m.Called(newAccount)
	return args.Get(0).(dao.Account)
}
