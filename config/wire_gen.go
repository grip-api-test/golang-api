// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"gin-gonic-api/app/controller"
	"gin-gonic-api/app/repository"
	"gin-gonic-api/app/service"
	"github.com/google/wire"
)

// Injectors from injector.go:

func Init() *Initialization {
	fakeAccountRepositoryImpl := repository.AccountRepositoryInit()
	localAccountServiceImpl := service.AccountServiceInit(fakeAccountRepositoryImpl)
	accountControllerImpl := controller.AccountControllerInit(localAccountServiceImpl)
	healthControllerImpl := controller.HealthControllerInit()
	initialization := NewInitialization(fakeAccountRepositoryImpl, localAccountServiceImpl, accountControllerImpl, healthControllerImpl)
	return initialization
}

// injector.go:

var accountServiceSet = wire.NewSet(service.AccountServiceInit, wire.Bind(new(service.AccountService), new(*service.LocalAccountServiceImpl)))

var accountRepoSet = wire.NewSet(repository.AccountRepositoryInit, wire.Bind(new(repository.AccountRepository), new(*repository.FakeAccountRepositoryImpl)))

var accountCtrlSet = wire.NewSet(controller.AccountControllerInit, wire.Bind(new(controller.AccountController), new(*controller.AccountControllerImpl)))

var healthCrtlSet = wire.NewSet(controller.HealthControllerInit, wire.Bind(new(controller.HealthController), new(*controller.HealthControllerImpl)))
