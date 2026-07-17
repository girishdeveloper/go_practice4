package main

import "fmt"

const SPEEDLIMIT = 100

func main() {
	var AtMainRoad float64 = SPEEDLIMIT
	var AtHighWay int = SPEEDLIMIT
	var AtRacingTrack string = string(SPEEDLIMIT) // ascii value
	fmt.Printf("%T is %3.2f\n", AtMainRoad, AtMainRoad)
	fmt.Printf("%T is %d\n", AtHighWay, AtHighWay)
	fmt.Printf("%T is %s\n", AtRacingTrack, AtRacingTrack)
}
