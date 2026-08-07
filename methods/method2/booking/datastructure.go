package booking

import (
	"Hotel/pricing"
	"Hotel/room"
	"time"
)

type BookingType int

const (
	AdvanceBooking BookingType = iota // 0
	SpotBooking                       // 1
	PartyBooking                      // 2
	BulkBooking                       // 3
)

var CutOff time.Time

type Booking struct {
	Type             BookingType
	TotalIndividuals int
	TotalBeds        int
	TotalRooms       int
	DurationInDays   int
	RoomType         room.RoomType
	RoomRate         pricing.PayRate
	BookedOn         time.Time
	GuestArrivalOn   time.Time
}

type Book struct {
	Orders     []Booking
	CutOffTime time.Time
}
