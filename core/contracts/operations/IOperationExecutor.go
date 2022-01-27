package operations

type IOperationExecutor interface {
	ExecuteOperation(string, []complex128) (complex128, error)
}
