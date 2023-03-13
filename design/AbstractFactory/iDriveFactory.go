package AbstractFactory

// ICarFactory 汽车工厂接口
type ICarFactory interface {
	MakeCar() ICar
	MakeBus() IBus
}

func GetCarFactory(brand string) ICarFactory {
	switch brand {
	case "Benz":
		return &Benz{}
	case "Audi":
		return &Audi{}
	default:
		return nil
	}

}
