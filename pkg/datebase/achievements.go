package datebase

import (
	"context"
	"github.com/beldmian/habitFarmer/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (db DB) GetAchievements() ([]models.Achievement, error) {
	client, err := db.Connect()
	if err != nil {
		return []models.Achievement{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("testing").Collection("achievements")
	cursor, err := collection.Find(ctx, bson.D{}, nil)
	if err != nil {
		return []models.Achievement{},err
	}
	ctx5, cancel5 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel5()
	var achievements []models.Achievement
	for cursor.Next(ctx5) {
		var achievement models.Achievement
		err := cursor.Decode(&achievement)
		if err != nil {
			return []models.Achievement{}, err
		}
		achievements = append(achievements, achievement)
	}
	if err := client.Disconnect(ctx); err != nil {
		return []models.Achievement{}, err
	}
	return achievements, nil
}
