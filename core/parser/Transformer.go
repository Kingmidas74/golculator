package parser

import (
	"errors"
	"golculator/core/collections"
	"golculator/core/contracts"
	icollections "golculator/core/contracts/collections"
	ilexemes "golculator/core/contracts/lexemes"
	ioperations "golculator/core/contracts/operations"
	"golculator/core/operations"
)

type Transformer struct {
	OperationList ioperations.IOperationList
}

func NewTransformer(operations ioperations.IOperationList) contracts.ITransformer {
	result := &Transformer{OperationList: operations}
	return result
}

func (this *Transformer) TransformToRPN(lexemes []ilexemes.ILexeme) (icollections.ICollection, error) {

	operationStack := collections.NewStack()
	result := collections.NewQueue()

	for i, lexeme := range lexemes {

		if lexeme.GetType() == ilexemes.DataLexeme {
			result.Push(lexeme)
			continue
		}

		if i == len(lexemes)-1 && lexeme.GetValue() != operations.CloseBracket {
			return nil, errors.New("operation without operand")
		}

		if lexeme.GetValue() == operations.OpenBracket || operationStack.Count() == 0 {
			operationStack.Push(lexeme)
			continue
		}
		if lexeme.GetValue() == operations.CloseBracket || lexeme.GetValue() == operations.Comma {
			for true {
				op, _ := operationStack.Pop()
				if op.GetValue() != operations.OpenBracket {
					result.Push(op)
				} else {
					if lexeme.GetValue() == operations.Comma {
						operationStack.Push(op)
					}
					break
				}
			}
			continue
		}

		currentOperation, err := this.OperationList.FindOperationByName(lexeme.GetValue())
		if err != nil {
			return nil, err
		}

		for operationStack.Count() > 0 {
			previousOperationLexeme, err := operationStack.Peek()
			if err != nil {
				return nil, err
			}
			previousOperation, err := this.OperationList.FindOperationByName(previousOperationLexeme.GetValue())
			if err != nil {
				return nil, err
			}
			if currentOperation.GetPriority() <= previousOperation.GetPriority() {
				op, _ := operationStack.Pop()
				result.Push(op)
				continue
			}
			break
		}
		operationStack.Push(lexeme)
	}

	for operationStack.Count() > 0 {
		op, _ := operationStack.Pop()
		result.Push(op)
	}

	return result, nil
}
