package repo

import (
	"context"
	"main/dtos"
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

	flight.RemainingSeats = flight.Seats
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

func (repo *FlightRepository) Update(flightAlter *dtos.FlightDto) error {
	client, cancel := utils.GetConn()
	defer cancel()

	idO, _ := primitive.ObjectIDFromHex(flightAlter.Id)

	filter := bson.M{"_id": idO}

	update := bson.M{"$set": bson.M{"flighdateandtime": flightAlter.FlighDateAndTime, "startingpoint": flightAlter.StartingPoint,
		"destination": flightAlter.Destination, "price": flightAlter.Price, "seats": flightAlter.Seats,
		"remainingseats": flightAlter.RemainingSeats}}

	coll := client.Database("airline").Collection("flights")
	_, err := coll.UpdateOne(context.TODO(), filter, update)

	return err
}

func (repo *FlightRepository) GetAll() ([]dtos.FlightDto, error) {
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

	var result []dtos.FlightDto
	result = ConvertToFlightDto(flights)

	return result, err
}

func (repo *FlightRepository) GetAllUpcoming() ([]dtos.FlightDto, error) {
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

	var result []dtos.FlightDto
	result = ConvertToFlightDto(flights)

	return result, err
}

func ConvertToFlightDto(data []model.Flight) []dtos.FlightDto {
	var result []dtos.FlightDto

	for _, flight := range data {
		var dtoFlight dtos.FlightDto
		dtoFlight.Id = flight.Id.Hex()
		dtoFlight.FlighDateAndTime = flight.FlighDateAndTime
		dtoFlight.StartingPoint = flight.StartingPoint
		dtoFlight.Destination = flight.Destination
		dtoFlight.Price = flight.Price
		dtoFlight.Seats = flight.Seats
		dtoFlight.RemainingSeats = flight.RemainingSeats

		result = append(result, dtoFlight)
	}

	return result
}