package repo

import (
	"context"

	"main/model"
	"main/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightRepository struct {
}

func (repo *FlightRepository) Create(flight *model.Flight) error {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("airline").Collection("flights")
	_, err := coll.InsertOne(context.TODO(), flight)

	return err
}

func (repo *FlightRepository) GetOne(id primitive.ObjectID) (model.Flight, error) {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("airline").Collection("flights")
	filter := bson.D{{Key: "_id", Value: id}}

	var result model.Flight
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}

func (repo *FlightRepository) Delete(id primitive.ObjectID) error {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("airline").Collection("flights")
	filter := bson.D{{Key: "_id", Value: id}}

	var result model.Flight
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err == nil {
		_, err2 := coll.DeleteOne(context.TODO(), result)
		return err2
	}

	return err
}

func (repo *FlightRepository) Update(flight *model.Flight) error {
	client, cancel := utils.GetConn()
	defer cancel()

	filter := bson.D{{Key: "_id", Value: flight.Id}}

	coll := client.Database("airline").Collection("flights")
	_, err := coll.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: flight}})

	return err
}

func (repo *FlightRepository) GetAll() ([]model.Flight, error) {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("airline").Collection("flights")

	var flights []model.Flight
	dataResult, err := coll.Find(context.TODO(), bson.M{})
	for dataResult.Next(context.TODO()) {
		var flight model.Flight
		err := dataResult.Decode(&flight)
		if err == nil {
			flights = append(flights, flight)
		}
	}

	return flights, err
}

func (repo *FlightRepository) GetAllUpcoming() ([]model.Flight, error) {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("airline").Collection("flights")
	filter := bson.M{"flighdateandtime": bson.M{"$gt": time.Now()}}

	var flights []model.Flight
	dataResult, err := coll.Find(context.TODO(), filter)
	for dataResult.Next(context.TODO()) {
		var flight model.Flight
		err := dataResult.Decode(&flight)
		if err == nil {
			flights = append(flights, flight)
		}
	}

	return flights, err
}
