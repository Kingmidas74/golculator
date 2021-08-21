package operations

type IOperation interface {
	GetName() string
	GetPriority() int
	GetArgumentsCount() int
	GetCode() string
	GetHandler() func([]float64)(float64,error)
}
