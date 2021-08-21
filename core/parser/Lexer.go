package parser

import (
	"errors"
	"golculator/core/contracts"
	"golculator/core/contracts/lexemes"
	"golculator/core/contracts/operations"
	"strconv"
)

type Lexer struct {
	OperationList operations.IOperationList
}

func NewLexer(operations operations.IOperationList) contracts.ILexer {
	result := &Lexer{OperationList: operations}
	return result
}

func(this *Lexer) Parse(expression string) ([]lexemes.ILexeme, error) {
	result := make([]lexemes.ILexeme,0)

	for i := 0; i < len(expression); i++ {
		currentNumber:=""
		for i<len(expression) {
			currentChar := string(expression[i])
			if _, err := strconv.Atoi(currentChar); err == nil || currentChar=="." {
				currentNumber+=currentChar
				i++
			} else {
				break
			}
		}

		if len(currentNumber)>0 {
			result = append(result, &Lexeme{
				Value: currentNumber,
				Type:  lexemes.DataLexeme,
			})
		}
		if i>=len(expression) {
			break
		}

		currentOperation := ""
		currentChar := string(expression[i])
		availableOperations := this.OperationList.FindOperationsStartsWith(currentChar)
		operationFound := false

		for _, availableOperation := range availableOperations {
			operationLength := len(availableOperation.GetName())
			possibleSignature := expression[i:(i+operationLength)]

			if possibleSignature == availableOperation.GetName() && ((availableOperation.GetPriority()==this.OperationList.GetMaxPriority() && string(expression[i+operationLength])=="(") || availableOperation.GetPriority()<this.OperationList.GetMaxPriority()) {
				currentOperation+=possibleSignature
				operationFound = true
				i+=operationLength-1
				break
			}
		}

		if operationFound {
			result = append(result, &Lexeme{
				Value: currentOperation,
				Type:  lexemes.OperationLexeme,
			})
		} else {
			return nil, errors.New("operation "+currentChar+" is wrong")
		}
	}

	return result, nil
}
