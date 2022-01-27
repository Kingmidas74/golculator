package helpers

type IArrayProvider interface {
	ReverseFloatArray(input []float64) []float64
	ReverseStringArray(input []string) []string
	ReverseComplexArray(input []complex128) []complex128
}

type ArrayProvider struct {
}

func NewArrayProvider() IArrayProvider {
	return &ArrayProvider{}
}

func (this *ArrayProvider) ReverseFloatArray(input []float64) []float64 {
	if len(input) == 0 {
		return input
	}
	return append(this.ReverseFloatArray(input[1:]), input[0])
}

func (this *ArrayProvider) ReverseStringArray(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(this.ReverseStringArray(input[1:]), input[0])
}

func (this *ArrayProvider) ReverseComplexArray(input []complex128) []complex128 {
	if len(input) == 0 {
		return input
	}
	return append(this.ReverseComplexArray(input[1:]), input[0])
}
