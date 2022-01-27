package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golculator/core"
	"golculator/core/helpers"
	"golculator/core/operations"
	"golculator/core/parser"
	"golculator/storage"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type OperationDTO struct {
	Name           string
	ArgumentsCount int
	Code           string
}

type WebServer struct {
	DB storage.Database
}

func NewWebServer(db storage.Database) WebServer {

	result := WebServer{DB: db}
	return result
}

func (this *WebServer) Run(staticPath, port string) {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File(fmt.Sprintf("%sindex.html", staticPath))
		} else {
			c.File(fmt.Sprintf("%s%s", staticPath, path.Join(dir, file)))
		}
	})

	r.POST("/operations", this.AddOperationHandler)
	r.GET("/operations", this.GetOperationsHandler)
	r.POST("/expressions", this.CalculateExpressionHandler)

	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}

func (this *WebServer) AddOperationHandler(context *gin.Context) {
	operationDto, statusCode, err := this.convertHTTPBodyToOperation(context.Request.Body)
	if err != nil {
		context.JSON(statusCode, err)
		return
	}
	if operationDto.Name == "+" || operationDto.Name == "-" || operationDto.Name == "*" || operationDto.Name == "/" {
		context.JSON(http.StatusBadRequest, "")
		return
	}

	this.DB.CreateOperation(storage.DBOperation{
		Name:           operationDto.Name,
		Priority:       3,
		ArgumentsCount: operationDto.ArgumentsCount,
		Code:           operationDto.Code,
	})
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
	}
	context.JSON(http.StatusCreated, gin.H{})
}

func (this *WebServer) GetOperationsHandler(context *gin.Context) {
	result := make([]OperationDTO, 0)
	availableOperations, err := this.DB.GetOperations()

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
	}
	for _, operation := range availableOperations {
		if operation.Priority > 0 {
			result = append(result, OperationDTO{
				Name:           operation.Name,
				ArgumentsCount: operation.ArgumentsCount,
				Code:           operation.Code,
			})
		}
	}
	context.JSON(http.StatusOK, result)
}

func (this *WebServer) CalculateExpressionHandler(context *gin.Context) {
	calculateRequest, statusCode, err := this.convertHTTPBodyToTodo(context.Request.Body)
	if err != nil {
		context.JSON(statusCode, err)
		return
	}

	availableOperationsDB, err := this.DB.GetOperations()

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
	}

	availableOperations := &operations.OperationList{}

	for _, op := range availableOperationsDB {
		availableOperations.Add(&operations.Operation{
			Name:           op.Name,
			Priority:       op.Priority,
			ArgumentsCount: op.ArgumentsCount,
			Code:           op.Code,
		})
	}

	actualLexer := parser.NewLexer(availableOperations)
	actualTransformer := parser.NewTransformer(availableOperations)
	actualOperationExecutor := operations.NewOperationExecutor(availableOperations)
	actualArrayProvider := helpers.NewArrayProvider()

	calculator := core.NewCalculator(actualLexer, actualTransformer, availableOperations, actualArrayProvider, actualOperationExecutor)

	result, err := calculator.Calculate(calculateRequest.Expression)
	if err != nil {
		context.JSON(500, err)
		return
	}
	context.JSON(statusCode, gin.H{"result": fmt.Sprintf("%G", result)})
}

func (this *WebServer) convertHTTPBodyToTodo(httpBody io.ReadCloser) (CalculateRequest, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return CalculateRequest{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return this.convertJSONBodyToTodo(body)
}

func (this *WebServer) convertJSONBodyToTodo(jsonBody []byte) (CalculateRequest, int, error) {
	var calculateRequest CalculateRequest
	err := json.Unmarshal(jsonBody, &calculateRequest)
	if err != nil {
		return CalculateRequest{}, http.StatusBadRequest, err
	}
	return calculateRequest, http.StatusOK, nil
}

func (this *WebServer) convertHTTPBodyToOperation(httpBody io.ReadCloser) (OperationDTO, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return OperationDTO{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return this.convertJSONBodyToOperation(body)
}

func (this *WebServer) convertJSONBodyToOperation(jsonBody []byte) (OperationDTO, int, error) {
	var result OperationDTO
	err := json.Unmarshal(jsonBody, &result)
	if err != nil {
		return OperationDTO{}, http.StatusBadRequest, err
	}
	return result, http.StatusOK, nil
}
