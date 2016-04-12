package main

import (
	"github.com/codegangsta/cli"
	"math/rand"
	"os"
	"time"
)

func main() {
	places := make([]string, 10, 15)
	places[0] = "Place 0"
	places[1] = "Veggie Fast Food"
	places[2] = "Chens"
	places[3] = "New Panda King"
	places[4] = "Al Madina's"
	places[5] = "Harveys"
	places[6] = "Williams"
	places[7] = "Subway"
	places[8] = "Mels"
	places[9] = "Mikeys"
	app := cli.NewApp()
	app.Name = "hungry"
	app.Usage = "Get a random restaurant"
	app.Action = func(c *cli.Context) {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		place := places[r1.Intn(9)]
		println(place)
	}
	app.Run(os.Args)
}
