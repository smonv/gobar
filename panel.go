package main

import (
	"fmt"

	"github.com/tthanh/gobar/block"
)

// Panel represent panel
type Panel struct {
	Left   map[string]*block.Base
	Center map[string]*block.Base
	Right  map[string]*block.Base
}

// NewPanel creat new panel
func NewPanel() *Panel {
	return &Panel{
		Left:   make(map[string]*block.Base),
		Center: make(map[string]*block.Base),
		Right:  make(map[string]*block.Base),
	}
}

// Build create panel message
func (p *Panel) Build(blocks map[string]*block.Base) string {
	bodies := []interface{}{}

	for _, b := range blocks {
		bodies = append(bodies, b.Text)
	}

	if len(bodies) == 0 {
		bodies = append(bodies, "")
	}

	return fmt.Sprintf("%v", bodies...)
}
