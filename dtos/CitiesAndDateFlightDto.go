package dtos

type CitiesAndDateFlightDto struct {
	Date          string `json:"date"`
	StartingPoint string `json:"startingPoint"`
	Destination   string `json:"destination"`
}
