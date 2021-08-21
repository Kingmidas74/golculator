package operations

import (
	"errors"
	"fmt"
	"golculator/core/contracts/operations"
	"strings"
)

type OperationList struct {
	operations []operations.IOperation
	maxPriority int
}

func NewOperationList() operations.IOperationList {
	result := &OperationList{make([]operations.IOperation,0),0}
	result.Add(&Operation{
		Name:           OpenBracket,
	})
	result.Add(&Operation{
		Name:           CloseBracket,
	})
	result.Add(&Operation{
		Name:           Comma,
	})
	return result
}

func(this *OperationList) Add(operation operations.IOperation) error {
	if op,_ := this.FindOperationByName(operation.GetName()); op != nil {
		return errors.New( fmt.Sprintf("operation %s already exsists", operation.GetName()))
	}
	this.operations = append(this.operations, operation)
	if this.maxPriority<operation.GetPriority() {
		this.maxPriority = operation.GetPriority()
	}
	return nil
}

func (this *OperationList) FindOperationByName(name string) (operations.IOperation,error) {
	for _, operation := range this.operations {
		if operation.GetName() == name {
			return operation,nil
		}
	}
	return nil,errors.New(fmt.Sprintf("operation %s not found", name))
}

func(this *OperationList) FindOperationsStartsWith(prefix string) []operations.IOperation {
	result := make([]operations.IOperation, 0)
	for _, operation := range this.operations {
		if strings.HasPrefix(operation.GetName(), prefix) {
			result = append(result, operation)
		}
	}
	return result
}

func(this *OperationList) GetMaxPriority() int {
	return this.maxPriority
}

func(this *OperationList) GetAll() []operations.IOperation {
	return this.operations
}


