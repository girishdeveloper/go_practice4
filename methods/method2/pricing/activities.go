package pricing

import (
	"Hotel/room"
	"fmt"
	"time"
)

func (pr *PayRates) Init() {
	pr.Rate = make([]PayRate, 0)
}

func (pr *PayRates) AddPayRate(rateOn PricingPer, duration int, price float64) {
	rate := PayRate{
		RateBy:   rateOn,
		Duration: duration,
		Price:    price,
	}
	pr.Rate = append(pr.Rate, rate)
}

func (pr *PayRates) ShowPayRates() {
	for _, v := range pr.Rate {
		fmt.Println(v)
	} // end of for
}

func (pr *PayRates) GetPayRate(index int) PayRate {
	if index < len(pr.Rate) {
		return pr.Rate[index]
	} else {
		return pr.Rate[0]
	}
}

func (tp TotalPayable) CalculatePayableAmount(rate PayRate, roomInfo room.RoomType,
	checkIn time.Time, checkOut time.Time, individuals int, beds int, rooms int) (amount float64) {
	var duration = float64(checkOut.Sub(checkIn).Hours() / 24)
	switch rate.RateBy {
	case 0:
		amount = float64(duration * float64(individuals) * rate.Price)
		break
	case 1:
		amount = float64(duration * float64(beds) * rate.Price)
		break
	case 2:
		amount = float64(duration * float64(rooms) * rate.Price)
		break
	case 3:
		amount = float64(duration * float64(rooms) * rate.Price)
		break
	} // end of switch
	return
}
