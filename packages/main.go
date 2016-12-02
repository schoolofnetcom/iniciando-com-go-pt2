package main

import (
	"packages/car"
	"fmt"
)

func main() {
	car := car.Car{"Gol", "Yellow"}
	fmt.Println(car.Start())
}
