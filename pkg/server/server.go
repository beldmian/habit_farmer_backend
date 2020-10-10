package server

import (
	"github.com/beldmian/habitFarmer/pkg/datebase"
	"github.com/gin-gonic/gin"
)

// Server ...
type Server struct {
	bindAddr string
	Router   *gin.Engine
	DB datebase.DB
}

func New(bindAddr string, DbUrl string) Server {
	r := gin.Default()
	db := datebase.DB{DbUrl: DbUrl}
	return Server{bindAddr,r, db}
}

func (s *Server) setupRouter() {
	s.Router.GET("/achievement", s.handleGetAchievements)

	s.Router.GET("/user/:email", s.handleGetUser)
	s.Router.POST("/user", s.handleCreateUser)
}

func (s Server) Start() error {
	s.setupRouter()
	return s.Router.Run(s.bindAddr)
}