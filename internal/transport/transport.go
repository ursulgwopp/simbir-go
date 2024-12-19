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
		}

		payment := api.Group("/Payment", t.userIdentity)
		{
			hesoyam := payment.Group("/Hesoyam")
			{
				hesoyam.POST("/:id", t.hesoyam)
			}
		}
	}

	return router
}
