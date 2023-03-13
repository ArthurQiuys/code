package AbstractFactory

type ICar interface {
	setDisplacement(Displacement int64)
	setModel(Model string)
	GetModel() string
	GetDisplacement() int64
}

type Car struct {
	model        string
	displacement int64
}

func (c *Car) setDisplacement(Displacement int64) {
	c.displacement = Displacement
}

func (c *Car) setModel(model string) {
	c.model = model
}

func (c *Car) GetModel() string {
	return c.model
}

func (c *Car) GetDisplacement() int64 {
	return c.displacement
}
