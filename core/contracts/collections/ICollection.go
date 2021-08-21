package collections

import "golculator/core/contracts/lexemes"

type ICollection interface {
	Push(value lexemes.ILexeme)
	Pop() (lexemes.ILexeme,error)
	Peek() (lexemes.ILexeme,error)

	Count() int
}