package collections

import (
	"errors"
	"golculator/core/contracts/collections"
	"golculator/core/contracts/lexemes"
	"sync"
)

type Queue struct {
	Head   collections.INode
	Tail   collections.INode
	Length int
	Lock *sync.Mutex
}

func NewQueue() collections.ICollection {
	result := &Queue{}
	result.Lock = &sync.Mutex{}
	return result
}

func (this *Queue) Count() int {
	this.Lock.Lock()
	defer this.Lock.Unlock()
	return this.Length
}

func (this *Queue) Push(value lexemes.ILexeme) {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	node := &Node{
		Value: value,
		Next:  nil,
	}

	if this.Tail == nil {
		this.Tail = node
		this.Head = node
	} else {
		this.Tail.SetNext(node)
		this.Tail = node
	}

	this.Length++
}

func (this *Queue) Pop() (lexemes.ILexeme,error) {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	if this.Head == nil {
		return nil, errors.New("node not found")
	}

	node := this.Head
	this.Head = node.GetNext()

	if this.Head == nil {
		this.Tail = nil
	}
	this.Length--

	return node.GetValue(), nil

}

func (this *Queue) Peek() (lexemes.ILexeme,error) {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	node := this.Head
	if node == nil || len(node.GetValue().GetValue()) == 0 {
		return nil, errors.New("node not found")
	}

	return node.GetValue(),nil
}

