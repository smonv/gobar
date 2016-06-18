package block

import (
	"sync"
	"time"

	"github.com/tthanh/gobar/message"
)

// DateBlock block
type DateBlock struct {
	Base
}

// NewDateBlock create new DateBlock
func NewDateBlock(name string, align string, bgColor string, fgColor string, interval int) *DateBlock {
	return &DateBlock{
		Base: Base{
			Name:     name,
			Align:    align,
			BgColor:  bgColor,
			FgColor:  fgColor,
			Interval: interval,
		},
	}
}

// Run implement block interface
func (d *DateBlock) Run(msgs chan message.Simple, stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Duration(d.Interval) * time.Second)
	for {
		select {
		case <-stop:
			return
		case <-ticker.C:
			t := time.Now().Format(time.RFC850)
			msg := message.Simple{
				Name:  d.Name,
				Align: d.Align,
				Text:  t,
			}
			msgs <- msg
		}
	}
}
