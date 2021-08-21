package operations

type IOperationExecutor interface {
	ExecuteOperation(string, []float64) (float64,error)
}
