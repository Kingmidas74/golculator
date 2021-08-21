package operations

type IOperation interface {
	GetName() string
	GetPriority() int
	GetArgumentsCount() int
	GetCode() string
}
