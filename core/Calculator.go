package core

import (
	"fmt"
	"golculator/core/collections"
	"golculator/core/contracts"
	ilexemes "golculator/core/contracts/lexemes"
	ioperations "golculator/core/contracts/operations"
	"golculator/core/helpers"
	"golculator/core/parser"
	"strconv"
)

type Calculator struct {
	CurrentLexer contracts.ILexer
	CurrentTransformer contracts.ITransformer
	CurrentOperations ioperations.IOperationList
	CurrentArrayProvider     helpers.IArrayProvider
	CurrentOperationExecutor ioperations.IOperationExecutor
}

func NewCalculator(currentLexer contracts.ILexer, currentTransformer contracts.ITransformer, currentOperations ioperations.IOperationList, currentArrayProvider helpers.IArrayProvider, currentOperationExecutor ioperations.IOperationExecutor) contracts.ICalculator {
	result := &Calculator{
		CurrentLexer:             currentLexer,
		CurrentTransformer:       currentTransformer,
		CurrentOperations:        currentOperations,
		CurrentArrayProvider:     currentArrayProvider,
		CurrentOperationExecutor: currentOperationExecutor,
	}
	return result
}

func(this *Calculator) Calculate(expression string) (float64,error) {

	lexemes, err := this.CurrentLexer.Parse(expression)
	if err != nil {
		return 0,err
	}

	chain, err := this.CurrentTransformer.TransformToRPN(lexemes)
	if err != nil {
		return 0, err
	}



	result := collections.NewStack()
	for chain.Count()>0 {
		c, err := chain.Pop()
		if err != nil {
			return 0, err
		}
		if c.GetType() == ilexemes.DataLexeme {
			result.Push(c)
			continue
		}
		op, err := this.CurrentOperations.FindOperationByName(c.GetValue())
		if err != nil {
			return 0, err
		}
		operands := make([]float64, 0)
		for i := 0; i < op.GetArgumentsCount() && result.Count()>0; i++ {
			a, err := result.Pop()
			if err != nil {
				return 0, err
			}
			parsedOperand, err := strconv.ParseFloat(a.GetValue(),64)
			if err != nil {
				return 0, err
			}
			operands = append(operands, parsedOperand)


			operands = this.CurrentArrayProvider.ReverseFloatArray(operands)
		}
		uresult, err := this.CurrentOperationExecutor.ExecuteOperation(op.GetName(), operands)
		if err != nil {
			return 0, err
		}

		result.Push(&parser.Lexeme{
			Value: fmt.Sprintf("%f",uresult),
			Type:  ilexemes.DataLexeme,
		})
	}
	unparsedResult,err := result.Pop()
	if err != nil {
		return 0, err
	}
	parsedResult, err := strconv.ParseFloat(unparsedResult.GetValue(),64)
	if err != nil {
		return 0, err
	}
	return parsedResult, nil
}

