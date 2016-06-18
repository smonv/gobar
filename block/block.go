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
	Text   = "%%{F%s}%%{B%s} %s %%{F-}%%{B-}"
)

// Base represent block base attributes
type Base struct {
	name     string
	align    string
	bgColor  string
	fgColor  string
	interval int
}

// Block interface
type Block interface {
	Run(msgs chan message.Simple, stop <-chan struct{}, wg *sync.WaitGroup)
	GetName() string
	GetAlign() string
}
