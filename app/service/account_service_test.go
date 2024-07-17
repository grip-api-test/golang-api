package service

import (
	"gin-gonic-api/app/domain/dao"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAll_FindsAllAccountsInRepository(t *testing.T) {
	mockAccountRepository := new(MockAccountRepository)
	sut := AccountServiceInit(mockAccountRepository)

	mockAccountRepository.On("FindAll").Return([]dao.Account{{Name: "Test Account"}})
	var accounts = sut.GetAll()

	assert.NotEmpty(t, accounts)
	assert.Contains(t, accounts, dao.Account{Name: "Test Account"})
	mockAccountRepository.AssertCalled(t, "FindAll")
}
