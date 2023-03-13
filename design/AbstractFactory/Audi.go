package AbstractFactory

type Audi struct {
}

// MakeBus 制造bus
func (b *Audi) MakeBus() IBus {
	return AudiBus{
		IBus: &Bus{
			model:        "2.0",
			displacement: 100,
		},
	}
}

// MakeCar 制造汽车
func (b *Audi) MakeCar() ICar {
	return AudiCar{
		ICar: &Car{
			model:        "1.0",
			displacement: 50,
		},
	}
}
