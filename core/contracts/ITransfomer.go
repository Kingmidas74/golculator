package contracts

import (
	"golculator/core/contracts/collections"
	"golculator/core/contracts/lexemes"
)

type ITransformer interface {
	TransformToRPN([]lexemes.ILexeme) (collections.ICollection,error)
}
