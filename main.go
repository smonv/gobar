package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/tthanh/gobar/block"
	"github.com/tthanh/gobar/message"
)

func main() {
	p := NewPanel()
	wg := new(sync.WaitGroup)

	stop := make(chan struct{})
	msgs := make(chan message.Simple)

	date := &block.Date{
		Base: block.Base{
			Name:     "datetime",
			Align:    block.Right,
			Interval: 1,
		},
	}

	wg.Add(1)
	go date.Run(msgs, stop, wg)

	go p.Start(msgs, stop)

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)

	<-osSignal
	close(stop)
	wg.Wait()
}
