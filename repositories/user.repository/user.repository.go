package user_repository

import (
	"Golang/database"
	m "Golang/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func Create(user m.User) error {

	var err error
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func Read() (m.Users, error) {

	var users m.Users

	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var user m.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func Update(user m.User, userID string) error {

	var err error
	uid, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": uid}
	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"password":   user.Password,
			"updated_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func Delete(userID string) error {

	var err error
	var uid primitive.ObjectID

	uid, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": uid}
	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
