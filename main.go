package main

import (
	"fmt"

	"github.com/SuttirakC/learn-basic-go/movie"
	"github.com/SuttirakC/learn-basic-go/ticket"
)

func init() {
	fmt.Println("init: main")
}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}