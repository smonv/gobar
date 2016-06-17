package block

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
	Text     string
	BgColor  string
	FgColor  string
	Interval int
}

// Block interface
type Block interface {
	Get(c chan Block)
}
