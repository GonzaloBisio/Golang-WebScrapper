package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Tipo de dato Usuario
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

// Lista de usuarios
type Users []*User
