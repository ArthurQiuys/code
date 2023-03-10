package Factory

// Car 对于结构体的定义，我们也可以将其放在单独的文件中，这样可以方便我们的管理
type Car struct {
	Name string
	Size int
}

// 为了方便我们的调用，我们可以为结构体定义一些方法
func (c *Car) setName(name string) {
	c.Name = name
}

func (c *Car) setSize(size int) {
	c.Size = size
}

func (c *Car) GetName() string {
	return c.Name
}

func (c *Car) GetSize() int {
	return c.Size
}
