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
