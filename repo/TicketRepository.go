package repo

import (
	"context"
	"errors"
	"main/model"
	"main/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketRepository struct {
}

func (repo *TicketRepository) Create(ticket *model.Ticket) error {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("airline").Collection("tickets")
	_, err := coll.InsertOne(context.TODO(), ticket)

	return err
}

func (repo *TicketRepository) GetOne(id primitive.ObjectID) (model.Ticket, error) {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("airline").Collection("tickets")
	filter := bson.D{{Key: "_id", Value: id}}

	var result model.Ticket
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}

func (repo *TicketRepository) GetAll(username string) ([]model.Ticket, error) {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("airline").Collection("tickets")
	filter := bson.D{{Key: "user", Value: username}}

	cursor, err := coll.Find(context.TODO(), filter)

	if err != nil {
		err = errors.New("Tickets for given User do not exist!")
		return []model.Ticket{}, err
	}

	var result []model.Ticket
	if err = cursor.All(context.TODO(), &result); err != nil {
		return []model.Ticket{}, err
	}

	return result, err
}

func (repo *TicketRepository) DeleteByFlight(flight primitive.ObjectID) error {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("airline").Collection("tickets")
	filter := bson.D{{Key: "flightid", Value: flight}}

	_ , err := coll.DeleteMany(context.Background(), filter)

	return err
}
