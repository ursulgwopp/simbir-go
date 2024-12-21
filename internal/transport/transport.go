package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-go/internal/models"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ursulgwopp/simbir-go/docs"
)

type Service interface {
	SignUp(req models.AccountRequest) (int, error)
	SignIn(req models.AccountRequest) (string, error)
	SignOut(token string) error

	GetAccount(accountId int) (models.AccountResponse, error)
	UpdateAccount(accountId int, req models.AccountRequest) error

	AdminListAccounts(from int, count int) ([]models.AdminAccountResponse, error)
	AdminGetAccount(accountId int) (models.AdminAccountResponse, error)
	AdminCreateAccount(req models.AdminAccountRequest) (int, error)
	AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error
	AdminDeleteAccount(accountId int) error

	Hesoyam(accountId int, userId int, isAdmin bool) error

	CreateTransport(ownerId int, req models.TransportRequest) (int, error)
	GetTransport(transportId int) (models.TransportResponse, error)
	UpdateTransport(userId int, transportId int, req models.TransportRequest) error
	DeleteTransport(userId int, transportId int) error

	AdminListTransports(from int, count int, transportType string) ([]models.AdminTransportResponse, error)
	AdminGetTransport(transportId int) (models.AdminTransportResponse, error)
	AdminCreateTransport(req models.AdminTransportRequest) (int, error)
	AdminUpdateTransport(transportId int, req models.AdminTransportRequest) error
	AdminDeleteTransport(transportId int) error

	CheckTokenIsValid(token string) (bool, error)
	ParseToken(token string) (models.TokenInfo, error)
}

type Transport struct {
	service Service
}

func NewTransport(service Service) *Transport {
	return &Transport{service: service}
}

func (t *Transport) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		account := api.Group("/Account")
		{
			account.POST("/SignUp", t.signUp)
			account.POST("/SignIn", t.signIn)
			account.POST("/SignOut", t.userIdentity, t.signOut)

			account.GET("/Me", t.userIdentity, t.me)
			account.PUT("/Update", t.userIdentity, t.update)
		}

		payment := api.Group("/Payment", t.userIdentity)
		{
			hesoyam := payment.Group("/Hesoyam")
			{
				hesoyam.POST("/:id", t.hesoyam)
			}
		}

		transport := api.Group("/Transport")
		{
			transport.GET("/:id", t.getTransport)
			transport.POST("/", t.userIdentity, t.createTransport)
			transport.PUT("/:id", t.userIdentity, t.updateTransport)
			transport.DELETE("/:id", t.userIdentity, t.deleteTransport)
		}

		admin := api.Group("/Admin")
		{
			account := admin.Group("/Account", t.adminIdentity)
			{
				account.GET("/", t.adminListAccounts)
				account.GET("/:id", t.adminGetAccount)
				account.POST("/", t.adminCreateAccount)
				account.PUT("/:id", t.adminUpdateAccount)
				account.DELETE("/:id", t.adminDeleteAccount)
			}

			transport := admin.Group("/Transport", t.adminIdentity)
			{
				transport.GET("/", t.adminListTransports)
				transport.GET("/:id", t.adminGetTransport)
				transport.POST("/", t.adminCreateTransport)
				transport.PUT("/:id", t.adminUpdateTransport)
				transport.DELETE("/:id", t.adminDeleteTransport)
			}
		}
	}

	return router
}
