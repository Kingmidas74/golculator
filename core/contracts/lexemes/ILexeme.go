package lexemes

type ILexeme interface {
	GetValue() string
	GetType() LexemeType
}
