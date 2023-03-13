package AbstractFactory

type Benz struct {
}

// MakeBus 制造bus
func (b *Benz) MakeBus() IBus {
	return BenzBus{
		IBus: &Bus{
			model:        "1.0",
			displacement: 50,
		},
	}
}

// MakeCar 制造car
func (b *Benz) MakeCar() ICar {
	return BenzCar{
		ICar: &Car{
			model:        "1.0",
			displacement: 150,
		},
	}
}
