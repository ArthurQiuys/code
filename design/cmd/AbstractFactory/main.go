package main

import (
	"code/design/AbstractFactory"
	"fmt"
)

func main() {
	Benz := AbstractFactory.GetCarFactory("Benz")
	Audi := AbstractFactory.GetCarFactory("Audi")
	BenzCar := Benz.MakeCar()
	BenzBus := Benz.MakeBus()
	AudiCar := Audi.MakeCar()
	AudiBus := Audi.MakeBus()
	printlnCar(BenzCar)
	printlnBus(BenzBus)
	printlnCar(AudiCar)
	printlnBus(AudiBus)
}

func printlnCar(i AbstractFactory.ICar) {
	fmt.Println(i.GetModel())
	fmt.Println(i.GetDisplacement())
}

func printlnBus(i AbstractFactory.IBus) {
	fmt.Println(i.GetModel())
	fmt.Println(i.GetDisplacement())
}
