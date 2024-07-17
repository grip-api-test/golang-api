package controller

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/domain/dao"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	AccountSummary(c *gin.Context)
	CreateAccount(c *gin.Context)
	AccountDetails(c *gin.Context)
}

type AccountControllerImpl struct {
	svc service.AccountService
}

func (u AccountControllerImpl) AccountSummary(c *gin.Context) {
	defer pkg.PanicHandler(c)
	accounts := u.svc.GetAll()
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, accounts))

}

func (u AccountControllerImpl) CreateAccount(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var request dao.Account
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}
	account := u.svc.CreateAccount(request)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, account))
}

func (u AccountControllerImpl) AccountDetails(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("accountID"))
	account := u.svc.Get(id)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, account))
}

func AccountControllerInit(accountService service.AccountService) *AccountControllerImpl {
	return &AccountControllerImpl{
		svc: accountService,
	}
}
