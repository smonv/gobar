package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"

	"github.com/tthanh/gobar/block"
	"github.com/tthanh/gobar/message"
)

func main() {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)

	p := NewPanel()
	wg := new(sync.WaitGroup)

	stop := make(chan struct{})
	msgs := make(chan message.Simple)
	bar := make(chan string)

	dateBlk := block.NewDateBlock("datetime", block.Right, "#FFd3d0c8", "#FF2d2d2d", 1)

	wg.Add(1)
	go dateBlk.Run(msgs, stop, wg)

	go p.Start(msgs, bar, stop)

	lemonBar := exec.Command("lemonbar",
		"-b",
		"-n", "lemonbar",
		"-g", "1366x20",
		"-f", "Source Sans Pro Bold:size=11",
		"-F", "#FFd3d0c8",
		"-B", "#FF2d2d2d")

	stdin, err := lemonBar.StdinPipe()
	if err != nil {
		panic(err)
	}

	lemonBar.Stdout = os.Stdout
	lemonBar.Stderr = os.Stderr

	if err = lemonBar.Start(); err != nil {
		panic(err)
	}

	for {
		select {
		case s := <-bar:
			io.Copy(stdin, bytes.NewBufferString(s))
			// io.Copy(os.Stdout, bytes.NewBufferString(s))
		case <-osSignal:
			err = lemonBar.Wait()
			if err != nil {
				fmt.Println(err)
			}
			close(stop)
			wg.Wait()
			return
		}
	}

}
