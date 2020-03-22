package main

import (
	"context"
	"fmt"
	"log"
	"math"

	gocorona "github.com/josephglaspie/gocorona"
)

func deathPercentage(confirmed, deaths int) float64 {
	x := float64(deaths) / float64(confirmed) * 100
	perc := math.Floor(float64(x)*100) / 100
	return perc
}

func main() {
	// client for accessing different endpoints of the API
	client := gocorona.Client{}
	ctx := context.Background()

	// GetLatestData returns total amonut confirmed cases, deaths, and recoveries.
	data, err := client.GetLatestData(ctx)
	if err != nil {
		log.Fatal("request failed:", err)
	}
	usa, err := client.GetDataByCountryCode(ctx, "US", false)
	if err != nil {
		log.Fatal("request failed:", err)
	}
	fmt.Println("##USA##")
	us := usa.Locations[0].Latest
	u := deathPercentage(us.Confirmed, us.Deaths)
	fmt.Printf("Cases: %d\nDeaths: %d\nRecoveries: %d\nDeath Rate: %v %%\n", us.Confirmed, us.Deaths, us.Recovered, u)
	fmt.Println("##GLOBAL##")
	p := deathPercentage(data.Data.Confirmed, data.Data.Deaths)
	fmt.Printf("Cases: %d\nDeaths: %d\nRecoveries: %d\nDeath Rate: %v %%\n", data.Data.Confirmed, data.Data.Deaths, data.Data.Recovered, p)
}
