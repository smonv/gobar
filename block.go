package main

type Block struct {
	Type     string
	Position string
	Body     string
}

func NewBlock(t string) *Block {
	return &Block{
		Type: t,
	}
}
