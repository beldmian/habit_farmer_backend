package server

import (
	"github.com/beldmian/habitFarmer/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) handleGetAchievements(c *gin.Context) {
	achievements, err := s.DB.GetAchievements()
	if err != nil {
		c.String(500, "%v",err.Error())
	}

	resp := getAchievementsResponse{OK: true, Achievements: achievements}

	c.JSON(200, resp)
}

func (s *Server) handleGetUser(c *gin.Context) {
	email := c.Param("email")
	user, err := s.DB.GetUser(email)
	if err != nil {
		c.String(500, "%v",err.Error())
	}

	c.JSON(200, user)
}

func (s Server) handleCreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	if err := s.DB.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(200, user)
}