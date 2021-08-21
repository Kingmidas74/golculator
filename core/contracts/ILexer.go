package contracts

import "golculator/core/contracts/lexemes"

type ILexer interface {
	Parse(string) ([]lexemes.ILexeme,error)
}
