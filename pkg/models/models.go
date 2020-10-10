package models

import "time"

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash" bson:"password_hash"`
	Habits []Habit `json:"habits"`
	Achievements []Achievement `json:"achievements"`
	Level int `json:"level"`
	XP int `json:"xp"`
}

type Habit struct {
	Dates []time.Time `json:"dates"`
	Name string `json:"name"`
}

type Achievement struct {
	Name string `json:"name"`
	Condition string `json:"condition"`
}