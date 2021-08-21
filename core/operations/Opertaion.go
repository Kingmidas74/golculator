package operations

type Operation struct {
	Name string
	Priority int
	ArgumentsCount int
	Code string
	Handler func([]float64)(float64,error)
}

func(this *Operation) GetName() string {
	return this.Name
}

func(this *Operation) GetPriority() int {
	return this.Priority
}

func(this *Operation) GetArgumentsCount() int {
	return this.ArgumentsCount
}

func(this *Operation) GetCode() string {
	return this.Code
}

func(this *Operation) GetHandler() func([]float64)(float64,error) {
	return this.Handler
}
