package main

import (
	"fmt"
	"golculator/core"
	ioperations "golculator/core/contracts/operations"
	"golculator/core/helpers"
	"golculator/core/operations"
	"golculator/core/parser"
	"log"
	"strconv"
	"strings"
)

func main() {

	input := "(2+2)*3*S1(2)+(1+POW(2,4)/6)"

	actualOperations,err := loadOperations("./operations")
	if err != nil {
		log.Fatal(err)
	}

	actualLexer := parser.NewLexer(actualOperations)
	actualTransformer := parser.NewTransformer(actualOperations)
	actualOperationExecutor := operations.NewOperationExecutor(actualOperations)
	actualArrayProvider := helpers.NewArrayProvider()

	calculator := core.NewCalculator(actualLexer,actualTransformer,actualOperations,actualArrayProvider,actualOperationExecutor)

	result,err := calculator.Calculate(input)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s = %f", input, result)

}

func loadOperations(source string) (ioperations.IOperationList,error) {

	operationsNameMap := map[string]string {
		"ADD": "+",
		"SUB": "-",
		"MUL": "*",
		"DIV": "/",
	}
	operationsPrioritiesMap := map[string]int {
		"ADD": 2,
		"SUB": 2,
		"MUL": 1,
		"DIV": 1,
	}

	actualOperations := operations.NewOperationList()

	fileReader := FileReader{}

	files,err := fileReader.GetListOfFiles(source)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		fileInfo := strings.Split(f.Name(),".")

		operationName := fileInfo[0]
		operationArgsCount,err := strconv.Atoi(fileInfo[1])
		if err != nil {
			return nil, err
		}
		operationCode, err := fileReader.ReadFileToString(fmt.Sprintf("%s/%s",source,f.Name()))
		if err != nil {
			return nil, err
		}

		if shortName, nameExist :=operationsNameMap[operationName]; nameExist {
			if priority, priorityExist := operationsPrioritiesMap[operationName]; priorityExist {
				actualOperations.Add(&operations.Operation{
					Name:           shortName,
					Priority:       priority,
					ArgumentsCount: operationArgsCount,
					Code:           operationCode,
				})
				continue
			}
		}
		actualOperations.Add(&operations.Operation{
			Name:           operationName,
			Priority:       3,
			ArgumentsCount: operationArgsCount,
			Code:           operationCode,
		})
	}

	return actualOperations, nil
}
