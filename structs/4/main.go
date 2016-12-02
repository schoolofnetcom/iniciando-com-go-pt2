package main

import (
	"fmt"
	"encoding/json"
)

type Car struct {
	Name  string `json:"Model"`
	Year  int `json:"-"`
	Color string
}

func main() {
	car := Car{"Gol", 2017, "Yellow"}

	result, _ := json.Marshal(car)
	//fmt.Println(result)
	fmt.Println(string(result))
}
