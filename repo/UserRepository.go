package repo

import (
	"context"
	"main/model"
	"main/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
}

func (repo *UserRepository) Create(user *model.User) error {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("airline").Collection("users")
	_, err := coll.InsertOne(context.TODO(), user)

	return err
}

func (repo *UserRepository) GetOne(username string) (model.User, error) {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("airline").Collection("users")
	filter := bson.D{ { Key: "username", Value: username } }

	var result model.User
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}
