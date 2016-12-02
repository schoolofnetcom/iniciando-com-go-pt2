package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	Name string
	CanFly bool
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\n Year: %d\n Color: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"Corolla", 2017, "White"}
	sCar := SuperCar{
		Car: car1,
		CanFly: true,
		Name: "Elantra",
	}
	fmt.Println(sCar.Name)
	fmt.Println(car1.info())

}
