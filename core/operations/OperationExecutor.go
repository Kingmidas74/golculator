package operations

import (
	lua "github.com/yuin/gopher-lua"
	ioperations "golculator/core/contracts/operations"
	luar "layeh.com/gopher-luar"
	"strconv"
)

type OperationExecutor struct {
	CurrentOperations ioperations.IOperationList
}

func NewOperationExecutor(operations ioperations.IOperationList) ioperations.IOperationExecutor {
	result := &OperationExecutor{CurrentOperations: operations}
	return result
}

func(this *OperationExecutor) ExecuteOperation(name string, arguments []float64) (float64,error) {

	operation, err := this.CurrentOperations.FindOperationByName(name)
	if err != nil {
		return 0, err
	}

	return this.callLua(operation,arguments)
}

func (this *OperationExecutor) callLua(operation ioperations.IOperation, arguments []float64) (float64,error) {

	L := lua.NewState()
	defer L.Close()

	args := &lua.LTable{}
	args.Metatable = lua.LNil

	for i, argument := range arguments {
		args.Insert(i,luar.New(L, argument))
	}

	resultLua := lua.LNumber(0)

	L.SetGlobal("args", args)
	L.SetGlobal("result", resultLua)

	if err := L.DoString(operation.GetCode()); err != nil {
		return 0, err
	}

	ret := L.Get(-1)
	return strconv.ParseFloat(ret.String(),64)
}
