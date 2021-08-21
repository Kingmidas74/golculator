package contracts

type ICalculator interface {
	Calculate(string) (float64,error)
}
