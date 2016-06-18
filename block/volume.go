package block

import (
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/tthanh/gobar/message"
	"github.com/tthanh/gobar/util"
)

// VolumeBlock represent block
type VolumeBlock struct {
	Base
}

// NewVolumeBlock create new VolumeBlock
func NewVolumeBlock(name string, align string, bgColor string, fgColor string, interval int) *VolumeBlock {
	return &VolumeBlock{
		Base: Base{
			name:     name,
			align:    align,
			bgColor:  bgColor,
			fgColor:  fgColor,
			interval: interval,
		},
	}
}

// Build create message
func (v *VolumeBlock) Build() message.Simple {
	// level, err := exec.Command("amixer", "get Master | grep 'Front Right:' | awk '{print $5}' | tr -d '[%]'").Output()
	levelAmixer := exec.Command("amixer", "get", "Master")
	levelGrep := exec.Command("grep", "'Front Right'")
	levelAwk := exec.Command("awk", "'{print $5}'")
	levelTr := exec.Command("tr", "-d", "'[%]'")

	level, err := util.PipeCommands(levelAmixer, levelGrep, levelAwk, levelTr)
	if err != nil {
		fmt.Println(err)
	}

	t := fmt.Sprintf(Text, v.fgColor, v.bgColor, level)

	return message.Simple{
		Name:  v.name,
		Align: v.align,
		Text:  t,
	}
}

// Run implement Block interface
func (v *VolumeBlock) Run(msgs chan message.Simple, stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(time.Duration(v.interval) * time.Second)
	for {
		select {
		case <-stop:
			return
		case <-ticker.C:
			msg := v.Build()
			msgs <- msg
		}
	}
}
