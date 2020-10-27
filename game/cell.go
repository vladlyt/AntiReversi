package game

const (
	BLACK = false
	WHITE = true
)

type Color bool

type Cell struct {
	IsFilled bool
	IsBlackHole bool
	Color    Color
}

func (c *Cell) Swap() {
	c.Color = !c.Color
}

func (c *Cell) Fill(color Color) {
	c.IsFilled = true
	c.Color = color
}

func (c Cell) String() string {
	if c.IsBlackHole {
		return "∅"
	}
	if !c.IsFilled {
		return " "
	}
	if c.Color {
		return "●"
	}
	return "○"
}
