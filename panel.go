package main

import (
	"fmt"

	"github.com/tthanh/gobar/block"
	"github.com/tthanh/gobar/message"
)

// Panel represent panel
type Panel struct {
	Left   map[string]message.Simple
	Center map[string]message.Simple
	Right  map[string]message.Simple
}

// NewPanel creat new panel
func NewPanel() *Panel {
	return &Panel{
		Left:   make(map[string]message.Simple),
		Center: make(map[string]message.Simple),
		Right:  make(map[string]message.Simple),
	}
}

// Build create panel message
func (p *Panel) Build(msgs map[string]message.Simple) string {
	bodies := []interface{}{}

	for _, m := range msgs {
		bodies = append(bodies, m.Text)
	}

	if len(bodies) == 0 {
		bodies = append(bodies, "")
	}

	return fmt.Sprintf("%v", bodies...)
}

// Start listening message
func (p *Panel) Start(msgs chan message.Simple, bar chan string, stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			return
		case msg := <-msgs:
			s := p.handleMessage(msg)
			bar <- s
		}
	}
}

func (p *Panel) handleMessage(m message.Simple) string {
	s := "%%{l}%s%%{c}%s%%{r}%s\n"

	pos := m.Align
	switch {
	case pos == block.Left:
		p.Left[m.Name] = m
	case pos == block.Center:
		p.Center[m.Name] = m
	case pos == block.Right:
		p.Right[m.Name] = m
	}

	l := p.Build(p.Left)
	c := p.Build(p.Center)
	r := p.Build(p.Right)

	return fmt.Sprintf(s, l, c, r)
}
