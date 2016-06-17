package main

import (
	"fmt"

	"github.com/tthanh/gobar/block"
)

func main() {
	panel := NewPanel()

	stopCh := make(chan struct{})
	fifoCh := make(chan block.Base)
	date := &block.Date{
		Base: block.Base{
			Name:     "datetime",
			Align:    block.Right,
			Interval: 1,
		},
	}
	go date.Get(fifoCh)

	for {
		select {
		case b := <-fifoCh:
			str := handleBlock(panel, &b)
			fmt.Println(str)
		case <-stopCh:
		}
	}
}

func handleBlock(panel *Panel, b *block.Base) string {
	p := "%%{l}%s%%{c}%s%%{r}%s"

	pos := b.Align
	switch {
	case pos == block.Left:
		panel.Left[b.Name] = b
	case pos == block.Center:

	case pos == block.Right:
		panel.Right[b.Name] = b
	}

	l := panel.Build(panel.Left)
	c := panel.Build(panel.Center)
	r := panel.Build(panel.Right)

	return fmt.Sprintf(p, l, c, r)
}
