package main

import (
	"Hotel/booking"
	"Hotel/guest"
	"Hotel/pricing"
	"Hotel/room"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func ShowMenu() {

	fmt.Println(`Welcome to the Hotel
		0) EXIT system
		1) Add property
		1.1) Add rooms
		1.2) Add property rates
		2) Check Availability
		3) Book room(s)
		3.1) Record guest detail
		3.2) Record booking detail
		3.3) Show bookings
		3.4) Cancel booking
		4) Guests
		4.1) Show guests
		4.2) Show guest detail
		4.3) Update detail
		4.4) Likely to visit`)
}

func ReturnToMenu() {
	move := bufio.NewReader(os.Stdin)
	fmt.Print("Press a key to move ahead")
	c, _, err := move.ReadRune()
	if err != nil {
		panic(err)
	}
	fmt.Print(c)
}

func main() {
	fmt.Println("Create a package and use it in main.go")
	log.Println("Create a module: go mod init <module-name>")
	log.Println(`import module packages in main.go: import (
	"<module-name>/<package-name1>"
	"<module-name>/<package-name2>"
	)`)
	var exitFlag bool = false
	var menuAction float32
	rs := &room.Rooms{}
	prs := &pricing.PayRates{}
	b := &booking.Book{}
	g := &guest.Customer{}
	a := &guest.Arrivals{}

	for exitFlag != true {
		ShowMenu()
		//take input from user
		fmt.Print("Which action do you want to proceed? ")
		fmt.Scanf("%f", &menuAction)
		switch menuAction {
		case float32(1.1):
			//add rooms
			rs.AddRoom(101, 1, 5, room.Dormitory, true)
			rs.AddRoom(102, 1, 5, room.Dormitory, false)
			rs.AddRoom(103, 1, 2, room.DoubleOccupancy, false)
			rs.AddRoom(201, 2, 2, room.DoubleOccupancy, true)
			rs.AddRoom(202, 2, 3, room.ResidentialSuite, false)
			rs.AddRoom(203, 2, 3, room.ResidentialSuite, false)
			rs.AddRoom(301, 3, 3, room.PresidentialSuite, false)
			log.Print("Rooms added!")
			break
		case float32(1.2):
			prs.AddPayRate(pricing.PerIndividual, 1, 30.45)
			prs.AddPayRate(pricing.PerBed, 1, 160.45)
			prs.AddPayRate(pricing.PerRoom, 1, 240.45)
			prs.AddPayRate(pricing.PerSuite, 1, 500.45)
			prs.ShowPayRates()
			break
		case float32(2):
			//check if rooms are available
			var fromDate time.Time = time.Now()
			//fmt.Printf("%T", fromDate)

			rs.RoomsRequired(fromDate, 3, 3)
			break
		case float32(3.1):
			var GuestIdx = g.AddGuest("PAN", "AUJEW1737Y", "Stebin",
				time.Date(1982, time.April, 13, 1, 20, 0, 0, room.Location), 9819356069,
				"Pot Pouri, Second Street, Gutta", "Kannur", "India", 340036)
			a.AddArrivalDetail(g.Guests[GuestIdx],
				time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+3, 0, 0, 0, 0, room.Location),
				3, guest.ByAir)
			log.Println("Guest detail added.")
			break
		case float32(3.2):
			after3Months3Days := time.Date(time.Now().Year(), time.Now().Month()+3, time.Now().Day()+3, 0, 0, 0, 0, room.Location)
			b.Book(booking.AdvanceBooking, 10, 10, 1, 2, room.Dormitory, prs.GetPayRate(0), time.Now(), after3Months3Days)
			//just now
			b.Book(booking.SpotBooking, 3, 3, 2, 2, room.DoubleOccupancy, prs.GetPayRate(1), time.Now(), time.Now())
			after3Days := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+3, 0, 0, 0, 0, room.Location)
			b.Book(booking.BulkBooking, 30, 15, 30, 2, room.ResidentialSuite, prs.GetPayRate(2), time.Now(), after3Days)
			after13Days := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+13, 0, 0, 0, 0, room.Location)
			b.Book(booking.PartyBooking, 5, 4, 1, 5, room.PresidentialSuite, prs.GetPayRate(3), time.Now(), after13Days)
			b.ShowBookings()
			break
		case float32(3.3):
			b.ShowBookings()
			break
		case float32(3.4):
			b.CancelBooking(0)
			b.ShowBookings()
			break
		case float32(4.1):
			g.ShowGuest()
			break
		case float32(0):
			exitFlag = true
			break
		} // end of switch

		if exitFlag != true {
			ReturnToMenu()
		}
	} // end of while
}
