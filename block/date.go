package block

import (
	"sync"
	"time"

	"github.com/tthanh/gobar/message"
)

// Date block
type Date struct {
	Base
}

// Run implement block interface
func (d *Date) Run(msgs chan message.Simple, stop <-chan struct{}, wg *sync.WaitGroup) {
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
