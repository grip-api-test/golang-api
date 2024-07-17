package service

import (
	"gin-gonic-api/app/domain/dao"

	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) FindAll() []dao.Account {
	args := m.Called()
	return args.Get(0).([]dao.Account)
}

func (m *MockAccountRepository) Save(account *dao.Account) dao.Account {
	args := m.Called(account)
	return args.Get(0).(dao.Account)
}

func (m *MockAccountRepository) Get(id int) dao.Account {
	args := m.Called(id)
	return args.Get(0).(dao.Account)
}
