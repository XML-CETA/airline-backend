package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"main/dtos"
	"main/model"
	"main/utils"
	"time"
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

func (repo *FlightRepository) SearchFlights(searchDto dtos.SearchDto) ([]model.Flight, error) {
	client, cancel := utils.GetConn()
	defer cancel()

	coll := client.Database("airline").Collection("flights")

	filter := createFilter(searchDto)

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	flights, err := searchByFilter(cursor, searchDto)

	return flights, err
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

func searchByFilter(cursor *mongo.Cursor, searchDto dtos.SearchDto) ([]model.Flight, error) {

	var flights []model.Flight
	for cursor.Next(context.TODO()) {
		var flight model.Flight
		err := cursor.Decode(&flight)

		if err != nil {
			return nil, err
		}
		if flight.RemainingSeats >= searchDto.NeededSeats {
			flights = append(flights, flight)
		}
	}

	return flights, nil
}

func createFilter(searchDto dtos.SearchDto) bson.D {
	startOfDay := time.Date(searchDto.FlighDateAndTime.Year(), searchDto.FlighDateAndTime.Month(), searchDto.FlighDateAndTime.Day(), 0, 0, 0, 0, searchDto.FlighDateAndTime.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1)
	filter := bson.D{
		{Key: "destination", Value: searchDto.Destination},
		{Key: "flighdateandtime", Value: bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		}},
		{Key: "startingpoint", Value: searchDto.StartingPoint},
	}
	return filter
}



