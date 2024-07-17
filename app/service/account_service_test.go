package service

import (
	"gin-gonic-api/app/domain/dao"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    shutdown()
    os.Exit(code)
} */

func Test_GetAll_FindsAllAccountsInRepository(t *testing.T) {
	mockAccountRepository := new(MockAccountRepository)
	sut := AccountServiceInit(mockAccountRepository)

	mockAccountRepository.On("FindAll").Return([]dao.Account{{Name: "Test Account"}})
	var accounts = sut.GetAll()

	assert.NotEmpty(t, accounts)
	assert.Contains(t, accounts, dao.Account{Name: "Test Account"})
	mockAccountRepository.AssertCalled(t, "FindAll")
}


func Test_GetAccountById(t *testing.T) {
	mockAccountRepository := new(MockAccountRepository)
	sut := AccountServiceInit(mockAccountRepository)
	mockAccountRepository.On("Get", 1).Return(dao.Account{Name: "Test Account"})
	
	var account = sut.Get(1)

	assert.NotNil(t, account)
	assert.Equal(t, "Test Account", account.Name)
	mockAccountRepository.AssertCalled(t, "Get", 1)
}