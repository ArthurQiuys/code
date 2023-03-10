package Factory

// 实现一个奔驰车的方法
type Benz struct {
	Car
}

// NewBenz 实现一个奔驰车的方法
func NewBenz() ICar {
	return &Benz{Car{Size: 5, Name: "Benz"}}
}
