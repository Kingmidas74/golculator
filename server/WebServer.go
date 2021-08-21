package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golculator/core/contracts"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
)

type CalculateRequest struct {
	Expression string `json:"expression"`
}

type WebServer struct {
	CurrentCalculator contracts.ICalculator
}

func NewWebServer(currentCalculator contracts.ICalculator) WebServer {
	result := WebServer{CurrentCalculator: currentCalculator}
	return result
}

func(this *WebServer) Run() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./ui/dist/ui/index.html")
		} else {
			c.File("./ui/dist/ui/" + path.Join(dir, file))
		}
	})

	r.POST("/operations", this.AddOperationHandler)
	r.GET("/operations", this.GetOperationsHandler)
	r.POST("/expressions", this.CalculateExpressionHandler)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}

func (this *WebServer) AddOperationHandler(context *gin.Context) {

}

func (this *WebServer) GetOperationsHandler(context *gin.Context) {

}

func (this *WebServer) CalculateExpressionHandler(context *gin.Context) {
	calculateRequest, statusCode, err := this.convertHTTPBodyToTodo(context.Request.Body)
	if err != nil {
		context.JSON(statusCode, err)
		return
	}
	result,err := this.CurrentCalculator.Calculate(calculateRequest.Expression)
	if err != nil {
		context.JSON(500, err)
		return
	}
	context.JSON(statusCode, gin.H{"result": result})
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