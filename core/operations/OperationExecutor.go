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

func (this *OperationExecutor) ExecuteOperation(name string, arguments []complex128) (complex128, error) {

	operation, err := this.CurrentOperations.FindOperationByName(name)
	if err != nil {
		return 0, err
	}

	return this.callLua(operation, arguments)
}

func (this *OperationExecutor) callLua(operation ioperations.IOperation, arguments []complex128) (complex128, error) {

	L := lua.NewState()
	defer L.Close()

	args := &lua.LTable{}
	args.Metatable = lua.LNil

	floatsArgs := make([]float64, 0)
	for _, argument := range arguments {
		floatsArgs = append(floatsArgs, real(argument))
		floatsArgs = append(floatsArgs, imag(argument))
	}

	for i, argument := range floatsArgs {
		args.Insert(i+1, luar.New(L, argument))
	}

	result := &lua.LTable{}
	result.Metatable = lua.LNil

	L.SetGlobal("args", args)
	L.SetGlobal("result", result)

	if err := L.DoString(operation.GetCode()); err != nil {
		return 0, err
	}

	realRet, _ := strconv.ParseFloat(result.RawGetInt(1).String(), 64)
	imagRet, _ := strconv.ParseFloat(result.RawGetInt(2).String(), 64)

	return complex(realRet, imagRet), nil
}
