package Factory

// ICar 对于接口的定义，我们可以将其放在单独的文件中，这样可以方便我们的管理
type ICar interface {
	setName(name string)
	setSize(size int)
	GetName() string
	GetSize() int
}
