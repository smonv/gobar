package block

import (
	"sync"

	"github.com/tthanh/gobar/message"
)

// Block Position
var (
	Left   = "left"
	Center = "center"
	Right  = "right"
)

// Base represent block base attributes
type Base struct {
	Name     string
	Align    string
	BgColor  string
	FgColor  string
	Interval int
}

// Block interface
type Block interface {
	Run(msgs chan message.Simple, stop <-chan struct{}, wg *sync.WaitGroup)
}
