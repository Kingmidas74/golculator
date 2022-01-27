package core

import (
	"golculator/core/helpers"
	"golculator/core/operations"
	"golculator/core/parser"
	"testing"
)

func TestCalculator_Calculate(t *testing.T) {
	availableOperations := &operations.OperationList{}

	_ = availableOperations.Add(&operations.Operation{
		Name:           "+",
		Priority:       1,
		ArgumentsCount: 2,
		Code: "result[1]=args[1]+args[3]\n" +
			"result[2]=args[2]+args[4]\n",
	})
	_ = availableOperations.Add(&operations.Operation{
		Name:           "-",
		Priority:       1,
		ArgumentsCount: 2,
		Code: "if args[3] == nil or args[3] == '' then\n" +
			"result[1]=args[1] * (-1)\n" +
			"result[2]=args[2] * (-1)\n" +
			"return\n" +
			"end\n" +
			"result[1]=args[1]-args[3]\n" +
			"result[2]=args[2]-args[4]\n",
	})
	_ = availableOperations.Add(&operations.Operation{
		Name:           "*",
		Priority:       2,
		ArgumentsCount: 2,
		Code: "result[1]=args[1]*args[3]-args[2]*args[4]\n" +
			"result[2]=args[1]*args[4]+args[2]*args[3]\n",
	})
	_ = availableOperations.Add(&operations.Operation{
		Name:           "/",
		Priority:       2,
		ArgumentsCount: 2,
		Code: "d=args[3]*args[3]+args[4]*args[4]\n" +
			"result[1]=(args[1]*args[3]+args[2]*args[4])/d\n" +
			"result[2]=(args[2]*args[3]-args[1]*args[4])/d",
	})
	_ = availableOperations.Add(&operations.Operation{
		Name:           "(",
		Priority:       0,
		ArgumentsCount: 0,
		Code:           "",
	})
	_ = availableOperations.Add(&operations.Operation{
		Name:           ")",
		Priority:       0,
		ArgumentsCount: 0,
		Code:           "",
	})
	_ = availableOperations.Add(&operations.Operation{
		Name:           ".",
		Priority:       0,
		ArgumentsCount: 0,
		Code:           "",
	})

	actualLexer := parser.NewLexer(availableOperations)
	actualTransformer := parser.NewTransformer(availableOperations)
	actualOperationExecutor := operations.NewOperationExecutor(availableOperations)
	actualArrayProvider := helpers.NewArrayProvider()

	calculator := NewCalculator(actualLexer, actualTransformer, availableOperations, actualArrayProvider, actualOperationExecutor)

	expressions := map[string]complex128{
		"1":                 1,
		"-1":                -1,
		"0":                 0,
		"1+1":               2,
		"1+0":               1,
		"(1)":               1,
		"(1+1)":             2,
		"((1)+(1))":         2,
		"1-1":               0,
		"0-1":               -1,
		"-1-1":              -2,
		"2*1":               2,
		"-1*2":              -2,
		"2*3":               6,
		"3*2":               6,
		"3+5i+1+12i":        4 + 17i,
		"3+5":               8,
		"3+5i":              3 + 5i,
		"0+3i+5+0i":         3i + 5,
		"(0+3i)+(0+5i)":     8i,
		"(2+0.4i)*(3-0.1i)": 6.04 + 1i,
		"(2+2i)/(4+6i)":     0.3846 + 0.076923077i,
	}

	comparatorProvider := helpers.NewComparatorProvider()

	for i, expression := range expressions {
		res, err := calculator.Calculate(i)
		if err != nil || !comparatorProvider.Equal(res, expression, 0.0001) {
			t.Errorf("%s =  %g; want %g", i, res, expression)
		}
	}

}
