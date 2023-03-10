package Factory

// Audi 实现一个奥迪车的方法
type Audi struct {
	Car
}

// NewAudi 实现一个奥迪车的方法
func NewAudi() ICar {
	return &Audi{Car{Size: 5, Name: "Audi"}}
}
