package guest

import "time"

type Guest struct {
	IdProof       string
	IdIdentifier  string
	Name          string
	DateOfBirth   time.Time
	ContactNumber int
	Address       string
	City          string
	Country       string
	PIN           int
}

type Customer struct {
	Guests []Guest
}

type TravelMode int

const (
	ByRoad   TravelMode = iota //0
	ByRail                     //1
	ByAir                      //2
	ByVessel                   //3
)

type Arrival struct {
	GuestDetail     Guest
	ArrivalOn       time.Time
	NoOfIndividuals int
	ArrivedBy       TravelMode
}

type Arrivals struct {
	Details []Arrival
}
