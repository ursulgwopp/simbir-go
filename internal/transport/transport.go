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

	Hesoyam(accountId int, userId int, isAdmin bool) error

	CreateTransport(ownerId int, req models.TransportRequest) (int, error)
	GetTransport(transportId int) (models.TransportResponse, error)
	UpdateTransport(userId int, transportId int, req models.TransportRequest) error
	DeleteTransport(userId int, transportId int) error

	GetAvailableTransport(latitude float64, longitude float64, radius float64, transportType string) ([]models.TransportResponse, error)
	GetRent(userId int, rentId int) (models.RentResponse, error)
	GetUserHistory(accountId int) ([]models.RentResponse, error)
	GetTransportHistory(userId int, transportId int) ([]models.RentResponse, error)
	StartRent(userId int, transportId int, rentType string) (int, error)
	StopRent(userId int, rentId int, latitude float64, longitude float64) error

	AdminListAccounts(from int, count int) ([]models.AdminAccountResponse, error)
	AdminGetAccount(accountId int) (models.AdminAccountResponse, error)
	AdminCreateAccount(req models.AdminAccountRequest) (int, error)
	AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error
	AdminDeleteAccount(accountId int) error

	AdminListTransports(from int, count int, transportType string) ([]models.AdminTransportResponse, error)
	AdminGetTransport(transportId int) (models.AdminTransportResponse, error)
	AdminCreateTransport(req models.AdminTransportRequest) (int, error)
	AdminUpdateTransport(transportId int, req models.AdminTransportRequest) error
	AdminDeleteTransport(transportId int) error

	AdminGetRent(rentId int) (models.RentResponse, error)
	AdminGetUserHistory(userId int) ([]models.RentResponse, error)
	AdminGetTransportHistory(transportId int) ([]models.RentResponse, error)
	AdminStartRent(req models.RentRequest) (int, error)
	AdminStopRent(rentId int, latitude float64, longitude float64) error
	AdminUpdateRent(rentId int, req models.RentRequest) error
	AdminDeleteRent(rentId int) error

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

		rent := api.Group("/Rent")
		{
			rent.GET("/Transport", t.getAvailableTransport)
			rent.GET("/MyHistory", t.userIdentity, t.getUserHistory)
			rent.GET("/TransportHistory/:id", t.userIdentity, t.getTransportHistory)
			rent.GET("/:id", t.userIdentity, t.getRent)
			rent.POST("/New/:id", t.userIdentity, t.startRent)
			rent.POST("/Stop/:id", t.userIdentity, t.stopRent)
		}

		admin := api.Group("/Admin", t.adminIdentity)
		{
			account := admin.Group("/Account")
			{
				account.GET("/", t.adminListAccounts)
				account.GET("/:id", t.adminGetAccount)
				account.POST("/", t.adminCreateAccount)
				account.PUT("/:id", t.adminUpdateAccount)
				account.DELETE("/:id", t.adminDeleteAccount)
			}

			transport := admin.Group("/Transport")
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
