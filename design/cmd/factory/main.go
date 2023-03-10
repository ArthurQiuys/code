package main

import (
	"code/design/Factory"
	"fmt"
)

func main() {
	var benz = Factory.MakeCar("Benz")
	var audi = Factory.MakeCar("Audi")
	fmt.Printf("benz%+v\n", benz)
	fmt.Printf("audi%+v\n", audi)
}
