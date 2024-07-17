package config

import (
	"gin-gonic-api/app/controller"
	"gin-gonic-api/app/repository"
	"gin-gonic-api/app/service"
)

type Initialization struct {
	accountRepo repository.AccountRepository
	accountSvc  service.AccountService
	AccountCtrl controller.AccountController
}

func NewInitialization(accountRepo repository.AccountRepository,
	accountService service.AccountService,
	accountCtrl controller.AccountController) *Initialization {
	return &Initialization{
		accountRepo: accountRepo,
		accountSvc:  accountService,
		AccountCtrl: accountCtrl,
	}
}
