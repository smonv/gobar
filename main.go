package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"syscall"
	"time"
)

func getDateTime(fifoCh chan Block) {
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	for {
		select {
		case <-ticker.C:
			b := Block{
				Type:     "datetime",
				Position: "right",
				Body:     time.Now().Format(time.RFC850),
			}
			fifoCh <- b
		}
	}
}

func main() {
	panel := NewPanel()
	xdg_runtime_dir := os.Getenv("XDG_RUNTIME_DIR")
	fifoPath := path.Join(xdg_runtime_dir, "panel_fifo")
	var start bool
	var stop bool

	flag.BoolVar(&start, "start", false, "start gobar")
	flag.BoolVar(&stop, "stop", false, "stop gobar")

	flag.Parse()

	syscall.Mkfifo(fifoPath, 0666)

	fifo, err := os.OpenFile(fifoPath, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	stopCh := make(chan struct{})
	fifoCh := make(chan Block)

	go getDateTime(fifoCh)

	for {
		select {
		case b := <-fifoCh:
			str := handleBlock(panel, &b)
			fifo.WriteString(str)
		case <-stopCh:
			fifo.Close()
			os.Remove(fifoPath)
		}
	}
}

func buildPanelOutput(blocks map[string]*Block) string {
	bodies := []interface{}{}

	for _, b := range blocks {
		bodies = append(bodies, b.Body)
	}

	if len(bodies) == 0 {
		bodies = append(bodies, "")
	}

	return fmt.Sprintf("%v", bodies...)
}

func handleBlock(panel *Panel, block *Block) string {
	p := "%%{l}%s%%{c}%s%%{r}%s\n"

	pos := block.Position
	switch {
	case pos == "left":
		panel.Left[block.Type] = block
	case pos == "center":

	case pos == "right":
		panel.Right[block.Type] = block
	}

	l := buildPanelOutput(panel.Left)
	c := buildPanelOutput(panel.Center)
	r := buildPanelOutput(panel.Right)

	return fmt.Sprintf(p, l, c, r)
}
