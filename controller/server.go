package controller

import (
	"final-project/config"
	"final-project/manager"
	"final-project/middleware"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager

	srv  *gin.Engine
	host string
}

func (s *server) Run() {
	// session
	store := cookie.NewStore([]byte("secret"))

	s.srv.Use(middleware.LoggerMiddleware())
	s.srv.Use(sessions.Sessions("session", store))

	// handler
	NewUserHandler(s.srv, s.usecaseManager.GetUserUsecase())
	NewLoginHandler(s.srv, s.usecaseManager.GetLoginUsecase())
	NewCustomerHandler(s.srv, s.usecaseManager.GetCustomerUsecase())
	NewMerchantHandler(s.srv, s.usecaseManager.GetMerchantUsecase())
	NewBankHandler(s.srv, s.usecaseManager.GetBankUsecase())
	NewTransferHandler(s.srv, s.usecaseManager.GetTransferUsecase())

	s.srv.Run(s.host)
}

func NewServer() Server {
	c := config.NewConfig()
	srv := gin.Default()
	infra := manager.NewInfraManager(c)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	if c.ApiHost == "" || c.ApiPort == "" {
		panic("No Host or port define")
	}

	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)
	return &server{
		usecaseManager: usecase,
		srv:            srv,
		host:           host,
	}
}
