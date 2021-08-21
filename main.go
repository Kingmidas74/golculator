package main

import (
	"fmt"
	"golculator/core"
	ioperations "golculator/core/contracts/operations"
	"golculator/core/helpers"
	"golculator/core/operations"
	"golculator/core/parser"
	"golculator/server"
	"log"
	"strconv"
	"strings"
)

func main() {

	actualOperations,err := loadOperations("./operations")
	if err != nil {
		log.Fatal(err)
	}

	actualLexer := parser.NewLexer(actualOperations)
	actualTransformer := parser.NewTransformer(actualOperations)
	actualOperationExecutor := operations.NewOperationExecutor(actualOperations)
	actualArrayProvider := helpers.NewArrayProvider()

	calculator := core.NewCalculator(actualLexer,actualTransformer,actualOperations,actualArrayProvider,actualOperationExecutor)

	webServer := server.NewWebServer(calculator, actualOperations)
	webServer.Run()
}

func loadOperations(source string) (ioperations.IOperationList,error) {

	operationsNameMap := map[string]string {
		"ADD": "+",
		"SUB": "-",
		"MUL": "*",
		"DIV": "/",
	}
	operationsPrioritiesMap := map[string]int {
		"ADD": 1,
		"SUB": 1,
		"MUL": 2,
		"DIV": 2,
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
