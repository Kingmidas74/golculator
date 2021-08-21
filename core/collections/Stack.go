package collections

import (
	"errors"
	"golculator/core/contracts/collections"
	"golculator/core/contracts/lexemes"
	"sync"
)

type Stack struct {
	Head   collections.INode
	Length int
	Lock *sync.Mutex
}

func NewStack() collections.ICollection {
	result := &Stack{}
	result.Lock = &sync.Mutex{}
	return result
}

func (this *Stack) Count() int {
	this.Lock.Lock()
	defer this.Lock.Unlock()
	return this.Length
}

func (this *Stack) Push(value lexemes.ILexeme) {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	node := &Node{
		Value: value,
		Next:  nil,
	}

	if this.Head == nil {
		this.Head = node
	} else {
		node.Next = this.Head
		this.Head = node
	}

	this.Length++
}

func (this *Stack) Pop() (lexemes.ILexeme,error) {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	var node collections.INode

	if this.Head != nil {
		node = this.Head
		this.Head = node.GetNext()
		this.Length--
	}

	if node == nil {
		return nil, errors.New("node not found")
	}

	return node.GetValue(), nil

}

func (this *Stack) Peek() (lexemes.ILexeme,error) {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	node := this.Head
	if node == nil || len(node.GetValue().GetValue()) == 0 {
		return nil, errors.New("node not found")
	}

	return node.GetValue(),nil
}

