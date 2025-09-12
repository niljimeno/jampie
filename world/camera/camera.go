package camera

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/niljimeno/jampie/utils/vector"
)

type Camera struct {
}

type DrawOptions struct {
	Position vector.Vector
	Size     vector.Vector
	Image    ebiten.Image
}

func (c *Camera) Update() {
}

func (c *Camera) Draw(d DrawOptions) {
}
