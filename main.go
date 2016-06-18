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

	blocks := []block.Block{}

	wg := new(sync.WaitGroup)

	stop := make(chan struct{})
	msgs := make(chan message.Simple)
	bar := make(chan string)

	volumeBlk := block.NewVolumeBlock("volume", block.Right, "#FFd3d0c8", "#FF2d2d2d", 5)
	blocks = append(blocks, volumeBlk)

	dateBlk := block.NewDateBlock("datetime", block.Right, "#FFd3d0c8", "#FF2d2d2d", 1)
	blocks = append(blocks, dateBlk)

	mPanel := NewPanel(blocks)

	for _, b := range blocks {
		wg.Add(1)
		go b.Run(msgs, stop, wg)
	}

	go run(mPanel, msgs, bar, stop)

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
		case <-osSignal:
			stdin.Close()
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

func run(mPanel *Panel, msgs chan message.Simple, bar chan string, stop <-chan struct{}) {
	lPanel := mPanel.Panels[block.Left]
	cPanel := mPanel.Panels[block.Center]
	rPanel := mPanel.Panels[block.Right]

	for {
		select {
		case <-stop:
			return
		case msg := <-msgs:
			switch msg.Align {
			case block.Left:
				lPanel.Build(msg)
			case block.Center:
				cPanel.Build(msg)
			case block.Right:
				rPanel.Build(msg)
			default:
			}

			s := fmt.Sprintf(mPanel.Pattern, lPanel.Text, cPanel.Text, rPanel.Text)
			bar <- s
		}
	}
}
