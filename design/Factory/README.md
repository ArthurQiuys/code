## 工厂模式

1. 产品接口
**此处定义了该产品的定义**
```go
type ICar interface {
setName(name string)
setSize(size int)
getName() string
getSize() int
}
```

2.实现对应的产品
**此处定义了该产品的实现**
```go
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
```

3.工厂接口
**此处定义了该工厂的定义**
```go
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
```
