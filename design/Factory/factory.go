package Factory

// MakeCar 制造一个汽车
func MakeCar(brand string) ICar {
	switch brand {
	case "Benz":
		return NewBenz()
	case "Audi":
		return NewAudi()
	default:
		return nil
	}
}
