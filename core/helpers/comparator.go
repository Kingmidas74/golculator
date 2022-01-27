package helpers

import "math"

type IComparatorProvider interface {
	Equal(a, b complex128, e float64) bool
}

type ComparatorProvider struct {
}

func NewComparatorProvider() IComparatorProvider {
	return &ComparatorProvider{}
}

func (this *ComparatorProvider) Equal(a, b complex128, e float64) bool {
	a1 := real(a)
	b1 := imag(a)
	a2 := real(b)
	b2 := imag(b)
	return math.Abs(math.Sqrt(a1*a1+b1*b1)-math.Sqrt(a2*a2+b2*b2)) < e
}
