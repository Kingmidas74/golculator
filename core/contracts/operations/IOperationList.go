package operations

type IOperationList interface {
	Add(IOperation) error
	FindOperationByName(string) (IOperation,error)
	FindOperationsStartsWith(string) []IOperation
	GetMaxPriority() int
	GetAll() []IOperation
}
