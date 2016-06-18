package block

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
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
	re := regexp.MustCompile("\\w+")

	amixer := exec.Command("amixer", "get", "Master")
	tail := exec.Command("tail", "-n", "1")

	output, _ := util.PipeCommands(amixer, tail)

	parts := strings.Fields(string(output))
	state := re.FindString(parts[len(parts)-1])
	level := re.FindString(parts[len(parts)-2])

	t := fmt.Sprintf(Text, v.fgColor, v.bgColor, state+" "+level)

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

// GetName implement Block interface
func (v *VolumeBlock) GetName() string {
	return v.name
}

// GetAlign implement Block interface
func (v *VolumeBlock) GetAlign() string {
	return v.align
}
