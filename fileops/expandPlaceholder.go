package main

import (
	"fmt"
	"os"
)

func main() {
	mapper := func(PlaceHolderName string) string {
		switch PlaceHolderName {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Girish"
		} // end of switch
		return ""
	}
	fmt.Println(os.Expand("Good ${DAY_PART}, ${NAME}!\n", mapper))
}
