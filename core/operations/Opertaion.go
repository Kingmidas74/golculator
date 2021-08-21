package operations

type Operation struct {
	Name string
	Priority int
	ArgumentsCount int
	Code string
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
