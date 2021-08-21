package collections

import "golculator/core/contracts/lexemes"

type INode interface {
	GetValue() lexemes.ILexeme
	SetValue(value lexemes.ILexeme)

	GetNext() INode
	SetNext(value INode)
}
