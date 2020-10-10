package datebase

import (
	"context"
	"github.com/beldmian/habitFarmer/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (db DB) GetUser(email string) (models.User, error) {
	client, err := db.Connect()
	if err != nil {
		return models.User{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database("testing").Collection("users")
	result := collection.FindOne(ctx, bson.D{{"email", email}})
	var user models.User
	if err := result.Decode(&user); err != nil {
		return models.User{}, nil
	}
	if err := client.Disconnect(ctx); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (db DB) CreateUser(user models.User) error {
	client, err := db.Connect()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database("testing").Collection("users")
	if _, err := collection.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}