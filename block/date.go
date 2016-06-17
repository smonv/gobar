package block

import (
	"time"
)

// Date block
type Date struct {
	Base
}

// Get implement block interface
func (d *Date) Get(c chan Base) {
	ticker := time.NewTicker(time.Duration(d.Interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			d.Text = time.Now().Format(time.RFC850)
			c <- d.Base
		}
	}
}
