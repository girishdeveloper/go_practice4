package room

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func (r *Rooms) Init() {
	r.rooms = make([]Room, 0)
}

func (r *Rooms) AddRoom(roomNumber uint32, floorNumber uint32, numberOfBeds uint8, roomType RoomType, availability bool) {
	room := Room{
		RoomNumber:  roomNumber,
		FloorNumber: floorNumber,
		Detail: RoomInfo{
			RoomNumber:   roomNumber,
			FloorNumber:  floorNumber,
			NumberOfBeds: numberOfBeds,
			Type:         roomType,
		},
		IsAvailable: availability,
	}
	r.rooms = append(r.rooms, room)
}

func (rm Room) RoomTypeString(roomType RoomType) string {
	var roomTypeString string
	switch roomType {
	case 0:
		roomTypeString = "Dormitory"
		break
	case 1:
		roomTypeString = "Single Occupancy"
		break
	case 2:
		roomTypeString = "Double Occupancy"
		break
	case 3:
		roomTypeString = "Residential Suite"
		break
	case 4:
		roomTypeString = "Presidential Suite"
		break
	default:
		roomTypeString = "Uncategorized  space"
	}
	return roomTypeString
}

func (r *Rooms) CheckAvailability(From time.Time, To time.Time, numberOfRooms uint32) (isAvailable bool, availableRooms uint32, roomDetails []RoomInfo) {
	availableRooms = 0
	isAvailable = false
	for _, v := range r.rooms {
		if v.IsAvailable == true {
			roomDetails = append(roomDetails, RoomInfo{
				RoomNumber:   v.RoomNumber,
				FloorNumber:  v.FloorNumber,
				NumberOfBeds: v.Detail.NumberOfBeds,
				Type:         v.Detail.Type,
			})
			availableRooms = availableRooms + 1
			if availableRooms == uint32(numberOfRooms) {
				break
			}
		}
	}
	if availableRooms == uint32(numberOfRooms) {
		isAvailable = true
	}
	return
}

func (r *Rooms) RoomsRequired(from time.Time, numberOfDays uint32, numberOfRooms uint32) {
	to := time.Date(from.Year(), from.Month(), from.Day()+int(numberOfDays), from.Hour(), from.Minute(), from.Second(), 0, Location)
	isRoomAvailable, roomsAvailable, roomDetails := r.CheckAvailability(from, to, numberOfRooms)
	rm := Room{}
	if isRoomAvailable == false {
		if roomsAvailable == 0 {
			log.Println("No rooms are available for the time period")
		} else if roomsAvailable < numberOfRooms {
			log.Print("Only ")
			log.Print(roomsAvailable)
			log.Print("rooms are available")
		}
	} else {
		log.Println(string(numberOfRooms) + "rooms are available")
	}
	if roomsAvailable > 0 {
		log.Println("Room Details are as follows: ")
		log.Println(roomDetails)
		fmt.Println(strings.Repeat("#*", 30))
		for i, v := range roomDetails {
			fmt.Println((i + 1), ") Room #: ", v.RoomNumber, ", Floor #: ", v.FloorNumber)
			fmt.Println("Occupancy: ", rm.RoomTypeString(v.Type), ", Beds:", v.NumberOfBeds)
		}
	}
}
