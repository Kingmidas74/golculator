package collections

import (
	"golculator/core/contracts/collections"
	"golculator/core/contracts/lexemes"
)

type Node struct {
	Value lexemes.ILexeme
	Next  collections.INode
}

func (this *Node) SetValue(value lexemes.ILexeme) {
	this.Value = value
}

func (this *Node) GetValue() lexemes.ILexeme {
	return this.Value
}

func (this *Node) SetNext(node collections.INode) {
	this.Next = node
}

func (this *Node) GetNext() collections.INode {
	return this.Next
}
