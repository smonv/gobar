package main

type Panel struct {
	Left   map[string]*Block
	Center map[string]*Block
	Right  map[string]*Block
}

func NewPanel() *Panel {
	return &Panel{
		Left:   make(map[string]*Block),
		Center: make(map[string]*Block),
		Right:  make(map[string]*Block),
	}
}
