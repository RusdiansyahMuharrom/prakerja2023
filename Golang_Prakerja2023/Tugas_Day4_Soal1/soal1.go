package main

import "fmt"

type Car struct {
	tipe   string
	fuelln float64
}

func (c Car) HitungJarak() float64 {
	return c.fuelln / 1.5
}

func main() {
	car := Car{"Avanza", 5}

	fmt.Println("Jarak Mobil", car.tipe, "=", car.HitungJarak(), "mil")
}
