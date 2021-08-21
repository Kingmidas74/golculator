package helpers

type IArrayProvider interface {
	ReverseFloatArray(input []float64) []float64
}

type ArrayProvider struct {

}

func NewArrayProvider() IArrayProvider {
	return &ArrayProvider{}
}

func(this *ArrayProvider) ReverseFloatArray(input []float64) []float64 {
	if len(input) == 0 {
		return input
	}
	return append(this.ReverseFloatArray(input[1:]), input[0])
}
