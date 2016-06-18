package main

import (
	"sync"

	"github.com/tthanh/gobar/block"
	"github.com/tthanh/gobar/message"
)

// Panel represent panel
type Panel struct {
	Panels  map[string]*Panel
	Blocks  map[string]string
	Keys    []string
	Text    string
	Pattern string
	sync.Mutex
}

// NewPanel creat new panel
func NewPanel(blocks []block.Block) *Panel {
	mPanel := &Panel{
		Panels:  make(map[string]*Panel),
		Pattern: "%%{l}%s%%{c}%s%%{r}%s\n",
	}

	lPanel := &Panel{
		Blocks: make(map[string]string),
	}
	cPanel := &Panel{
		Blocks: make(map[string]string),
	}
	rPanel := &Panel{
		Blocks: make(map[string]string),
	}

	for _, b := range blocks {
		switch b.GetAlign() {
		case block.Left:
			lPanel.Blocks[b.GetName()] = ""
			lPanel.Keys = append(lPanel.Keys, b.GetName())
		case block.Center:
			cPanel.Blocks[b.GetName()] = ""
			cPanel.Keys = append(cPanel.Keys, b.GetName())
		case block.Right:
			rPanel.Blocks[b.GetName()] = ""
			rPanel.Keys = append(rPanel.Keys, b.GetName())
		default:
		}
	}

	mPanel.Panels[block.Left] = lPanel
	mPanel.Panels[block.Center] = cPanel
	mPanel.Panels[block.Right] = rPanel

	return mPanel
}

// Build create panel text
func (p *Panel) Build(msg message.Simple) {
	p.Lock()
	defer p.Unlock()

	p.Blocks[msg.Name] = msg.Text
	pStr := ""
	for _, k := range p.Keys {
		s := p.Blocks[k]
		pStr += s
	}
	p.Text = pStr
}
