package server

import (
	"github.com/beldmian/habitFarmer/pkg/models"
)

type getAchievementsResponse struct {
	OK bool `json:"ok"`
	Achievements []models.Achievement `json:"achievements"`
}
