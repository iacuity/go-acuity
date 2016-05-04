package solitary

type Container interface {
	Size() uint
	Insert(key string) bool
	Delete(key string) bool
	Exists(key string) bool
}

type ContainerType int

const (
	CuckooFilterType ContainerType = iota
)

// return the new container
func NewContainer(capacity uint, cntnrType ContainerType) Container {
	var container Container
	switch cntnrType {
	case CuckooFilterType:
		container = newCuckooContainer(capacity)
	default:
		container = newCuckooContainer(capacity)
	}

	return container
}
