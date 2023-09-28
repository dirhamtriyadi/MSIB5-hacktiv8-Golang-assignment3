package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Weather struct {
	Wind  int
	Water int
}

func main() {
	for {
		time.Sleep(15 * time.Second)
		randomValueForWind := rand.Intn(10)
		randomValueForWater := rand.Intn(10)

		var weather Weather = Weather{
			Wind:  randomValueForWind,
			Water: randomValueForWater,
		}

		a, _ := json.Marshal(weather)

		fmt.Printf("%s\n", a)

		if randomValueForWater < 5 {
			fmt.Println("status water : aman")
		} else if randomValueForWater > 5 && randomValueForWater < 9 {
			fmt.Println("status water : siaga")
		} else if randomValueForWater > 9 {
			fmt.Println("status water : bahaya")
		}

		if randomValueForWind < 7 {
			fmt.Println("status wind : aman")
		} else if randomValueForWind > 6 && randomValueForWind < 16 {
			fmt.Println("status wind : siaga")
		} else if randomValueForWind > 15 {
			fmt.Println("status wind : bahaya")
		}
		fmt.Println("=====================================")
	}
}
