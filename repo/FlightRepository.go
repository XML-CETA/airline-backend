package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

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

func (repo *FlightRepository) SearchFlights(searchDto dtos.SearchDto) ([]dtos.SearchedFlightDto, error) {
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

	var result []dtos.SearchedFlightDto
	result = ConvertToSearchedFlightDto(flights, searchDto.NeededSeats)

	return result, nil
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

func ConvertToSearchedFlightDto(data []model.Flight, neededSeats int) []dtos.SearchedFlightDto {
	var result []dtos.SearchedFlightDto

	for _, flight := range data {
		var searchedFlightDto dtos.SearchedFlightDto
		searchedFlightDto.Id = flight.Id.Hex()
		searchedFlightDto.FlighDateAndTime = flight.FlighDateAndTime
		searchedFlightDto.StartingPoint = flight.StartingPoint
		searchedFlightDto.Destination = flight.Destination
		searchedFlightDto.Price = flight.Price
		searchedFlightDto.TotalPrice = searchedFlightDto.Price * neededSeats
		searchedFlightDto.NeededSeats = neededSeats

		result = append(result, searchedFlightDto)
	}

	return result
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
