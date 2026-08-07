package room

import (
	"time"
)

var Location, _ = time.LoadLocation("Asia/Kolkata")

type RoomType int

const (
	Dormitory         RoomType = iota //0
	SingleOccupancy                   //1
	DoubleOccupancy                   //2
	ResidentialSuite                  //3
	PresidentialSuite                 //4
)

type RoomInfo struct {
	RoomNumber   uint32
	FloorNumber  uint32
	NumberOfBeds uint8
	Type         RoomType
}

type Room struct {
	RoomNumber  uint32
	FloorNumber uint32
	Detail      RoomInfo
	IsAvailable bool
}

type Rooms struct {
	rooms []Room
}
