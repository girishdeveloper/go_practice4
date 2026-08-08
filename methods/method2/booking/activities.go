package booking

import (
	"Hotel/pricing"
	"Hotel/room"
	"fmt"
	"slices"
	"strings"
	"time"
)

func (b *Book) Init() {
	b.Orders = make([]Booking, 0)
	b.CutOffTime = b.GetCutOff()
}

func (b *Book) GetCutOff() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, room.Location)
}

func (b *Book) Book(bookType BookingType, individuals int, beds int, rooms int, duration int,
	roomType room.RoomType, roomRate pricing.PayRate, bookingTime time.Time, arrivalTime time.Time) {
	if b.GetCutOff().Sub(arrivalTime) > 0 {
		fmt.Println("Booking should be done atleast a day before")
	} else {
		var order = Booking{
			Type:             bookType,
			TotalIndividuals: individuals,
			TotalBeds:        beds,
			TotalRooms:       rooms,
			DurationInDays:   duration,
			RoomType:         roomType,
			RoomRate:         roomRate,
			BookedOn:         bookingTime,
			GuestArrivalOn:   arrivalTime,
		}
		b.Orders = append(b.Orders, order)
	}
}

func (b *Book) ShowBookingType(intType BookingType) string {
	var typeString string = ""
	switch intType {
	case 0:
		typeString = "Advance Booking"
		break
	case 1:
		typeString = "Spot Booking"
		break
	case 2:
		typeString = "Party Booking"
		break
	case 3:
		typeString = "Bulk Booking"
		break
	} // end of for
	return typeString
}

func (b *Book) ShowBookings() {
	rm := room.Room{}
	tp := pricing.TotalPayable{}
	for _, v := range b.Orders {
		checkOutOn := time.Date(v.GuestArrivalOn.Year(), v.GuestArrivalOn.Month(), v.GuestArrivalOn.Day()+v.DurationInDays, 0, 0, 0, 0, room.Location)
		fmt.Println("check out date", checkOutOn)
		amount := tp.CalculatePayableAmount(v.RoomRate, v.RoomType, v.GuestArrivalOn, checkOutOn, v.TotalIndividuals, v.TotalBeds, v.TotalRooms)

		fmt.Println("Booking Type: " + b.ShowBookingType(v.Type))
		fmt.Printf("Individuals: %d persons \t Beds: %d\n", v.TotalIndividuals, v.TotalBeds)
		fmt.Printf("Rooms: %d \t Duration: %d days\n", v.TotalRooms, v.DurationInDays)
		fmt.Printf("Room type: %s \t Payable amount: %0.2f\n", rm.RoomTypeString(v.RoomType), amount)
		fmt.Printf("Booked on: %v \t Guest will arrive on: %v\n", v.BookedOn.Format(time.DateTime), v.GuestArrivalOn.Format(time.DateTime))
		fmt.Println(strings.Repeat("**", 15))
	}
}

func (b *Book) CancelBooking(index int) {
	if len(b.Orders) >= 1 && len(b.Orders) > index {
		b.Orders = slices.Delete(b.Orders, index, 1)
	} else {
		fmt.Println("No bookings found for deletion")
	}
}
