package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	ioperations "golculator/core/contracts/operations"
	"golculator/core/operations"
	"golculator/server"
	"golculator/storage"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	currentEnvironment := os.Getenv("ENVIRONMENT")
	if len(currentEnvironment) == 0 || currentEnvironment == "Development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := GetDB()

	if err != nil {
		log.Fatal(err)
	}

	argsWithoutProg := os.Args[1:]

	switch argsWithoutProg[0] {
	case "seed":
		{
			err = Seeding(db)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	webServer := server.NewWebServer(db)
	webServer.Run("./static/", os.Getenv("APP_PORT"))
}

func GetDB() (storage.Database, error) {

	db := storage.Database{}

	if err := db.Initialize(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")); err != nil {
		return storage.Database{}, err
	}

	return db, nil
}

func Seeding(db storage.Database) error {
	actualOperations, err := loadOperations("./operations/")
	if err != nil {
		return err
	}

	err = db.Up()
	if err != nil {
		return err
	}

	for _, operation := range actualOperations.GetAll() {
		db.CreateOperation(storage.DBOperation{
			Name:           operation.GetName(),
			ArgumentsCount: operation.GetArgumentsCount(),
			Priority:       operation.GetPriority(),
			Code:           operation.GetCode(),
		})
	}
	return nil
}

func loadOperations(source string) (ioperations.IOperationList, error) {

	operationsNameMap := map[string]string{
		"ADD": "+",
		"SUB": "-",
		"MUL": "*",
		"DIV": "/",
	}
	operationsPrioritiesMap := map[string]int{
		"ADD": 1,
		"SUB": 1,
		"MUL": 2,
		"DIV": 2,
	}

	actualOperations := operations.NewOperationList()

	fileReader := FileReader{}

	files, err := fileReader.GetListOfFiles(source)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		fileInfo := strings.Split(f.Name(), ".")

		operationName := fileInfo[0]
		operationArgsCount, err := strconv.Atoi(fileInfo[1])
		if err != nil {
			return nil, err
		}
		operationCode, err := fileReader.ReadFileToString(fmt.Sprintf("%s%s", source, f.Name()))
		if err != nil {
			return nil, err
		}

		if shortName, nameExist := operationsNameMap[operationName]; nameExist {
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
