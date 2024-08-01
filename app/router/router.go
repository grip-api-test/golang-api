package router

import (
	"gin-gonic-api/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		accounts := api.Group("/accounts")
		{
			accounts.GET("", init.AccountCtrl.AccountSummary)
			accounts.POST("", init.AccountCtrl.CreateAccount)
			accounts.GET("/:accountID", init.AccountCtrl.AccountDetails)
		}
		health := api.Group("/health")
		{
			health.GET("", init.HealthCtrl.Health)
		}
	}

	return router
}
