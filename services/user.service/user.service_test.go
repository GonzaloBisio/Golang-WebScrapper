package user_service_test

import (
	m "Golang/models"
	userService "Golang/services/user.service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

var userID string

func TestCreate(t *testing.T) {

	oid := primitive.NewObjectID()
	userID = oid.Hex()

	user := m.User{
		ID:        oid,
		Name:      "John Doe",
		Email:     "john.doe@gmail.com",
		Password:  "123456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := userService.Create(user)
	if err != nil {
		t.Error("User create test failed")
		t.Fail()
	} else {
		t.Log("User create test passed")
	}

}

func TestRead(t *testing.T) {

	users, err := userService.Read()
	if err != nil {
		t.Error("User read test failed")
		t.Fail()
	}
	if len(users) == 0 {
		t.Error("La consulta no retornó datos")
		t.Fail()
	} else {
		t.Log("User read test passed")
	}

}

func TestUpdate(t *testing.T) {
	user := m.User{
		Name:     "Jesus",
		Email:    "jesus@gmail.com",
		Password: "updatedPass",
	}
	err := userService.Update(user, userID)

	if err != nil {
		t.Error("Update test failed")
		t.Fail()
	} else {
		t.Log("Update test succeded!")
	}
}

func TestDelete(t *testing.T) {
	err := userService.Delete(userID)
	if err != nil {
		t.Error("Delete test failed")
		t.Fail()
	} else {
		t.Log("Delete test passed!")
	}

}
