package AbstractFactory

type IBus interface {
	setDisplacement(Displacement int64)
	setModel(Model string)
	GetModel() string
	GetDisplacement() int64
}

type Bus struct {
	model        string
	displacement int64
}

func (b *Bus) setDisplacement(Displacement int64) {
	b.displacement = Displacement
}

func (b *Bus) setModel(model string) {
	b.model = model
}

func (b *Bus) GetModel() string {
	return b.model
}

func (b *Bus) GetDisplacement() int64 {
	return b.displacement
}
