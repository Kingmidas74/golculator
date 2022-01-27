package contracts

type ICalculator interface {
	Calculate(string) (complex128, error)
}
