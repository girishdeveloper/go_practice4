package guest

import (
	"fmt"
	"time"
)

func (g *Customer) Init() {
	g.Guests = make([]Guest, 0)
}

func (a *Arrivals) Init() {
	a.Details = make([]Arrival, 0)
}

func (g *Customer) AddGuest(idProof string, idIdentifier string, name string, dob time.Time,
	contactNumber int, address string, city string, country string, pincode int) int {
	g.Guests = append(g.Guests, Guest{
		IdProof:       idProof,
		IdIdentifier:  idIdentifier,
		Name:          name,
		DateOfBirth:   dob,
		ContactNumber: contactNumber,
		Address:       address,
		City:          city,
		Country:       country,
		PIN:           pincode,
	})
	return len(g.Guests) - 1
}

func (g *Customer) ShowGuest() {
	fmt.Println(g.Guests)
}

func (a *Arrivals) AddArrivalDetail(guest Guest, arrivalOn time.Time, totalIndividuals int,
	arrivedBy TravelMode) {
	a.Details = append(a.Details, Arrival{
		GuestDetail:     guest,
		ArrivalOn:       arrivalOn,
		NoOfIndividuals: totalIndividuals,
		ArrivedBy:       arrivedBy,
	})
}
