package repository

import (
	"gin-gonic-api/app/domain/dao"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SaveAccount_AddsItToTheRepository(t *testing.T) {
	account := dao.Account{Name: "Test"}
	sut := AccountRepositoryInit()

	sut.Save(&account)

	var accounts = sut.FindAll()
	assert.NotEmpty(t, accounts)
}


func Test_GetAccountByID(t *testing.T) {
	id := 1;
	sut := AccountRepositoryInit()
	sut.Save(&dao.Account{ID: 1, Name: "Denis"})
	
	newAccount := sut.Get(id)

	assert.Equal(t, 1, newAccount.ID)
	assert.Equal(t, "Denis", newAccount.Name)
}