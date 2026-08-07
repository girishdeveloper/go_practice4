package pricing

import "time"

type PricingPer int

const (
	PerIndividual PricingPer = iota //0
	PerBed                          //1
	PerRoom                         //2
	PerSuite                        //3
)

type CheckInTime time.Time
type CheckOutTime time.Time

type PayRate struct {
	RateBy   PricingPer
	Duration int
	Price    float64
}

type PayRates struct {
	Rate []PayRate
}

type TotalPayable struct {
	PaymentRate      PayRate
	CheckIn          CheckInTime
	CheckOut         CheckOutTime
	TotalIndividuals int
	TotalRooms       int
	Total            float64
}
