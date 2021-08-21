package parser

import "golculator/core/contracts/lexemes"

type Lexeme struct {
	Value string
	Type lexemes.LexemeType
}

func(this *Lexeme) GetValue() string {
	return this.Value
}

func(this *Lexeme) GetType() lexemes.LexemeType {
	return this.Type
}
