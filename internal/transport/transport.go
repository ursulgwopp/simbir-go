package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-go/internal/models"
)

type Service interface {
	SignUp(req models.AuthRequest) (int, error)
	SignIn(req models.AuthRequest) (string, error)
	SignOut(token string) error

	GetAccount(accountId int) (models.AccountResponse, error)
	UpdateAccount(accountId int, req models.AccountResponse)
}

type Transport struct {
	service Service
}

func NewTransport(service Service) *Transport {
	return &Transport{service: service}
}

func (t *Transport) InitRoutes() *gin.Engine {
	router := gin.Default()

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		account := api.Group("/Account")
		{
			account.POST("/SignUp", nil)
			account.POST("/SignIn", nil)
			account.POST("/SignOut", nil)

			account.GET("/Me", nil)
			account.PUT("/Update", nil)
		}
	}

	return router
}
