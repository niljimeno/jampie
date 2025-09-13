package camera

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/niljimeno/jampie/utils/vector"
)

type Camera struct {
	Screen   *ebiten.Image
	TargetY  float64
	Position vector.Vector

	SpeedReduction float64
	Margin         float64
}

type DrawOptions struct {
	Position vector.Vector
	Size     vector.Vector
	Image    *ebiten.Image
	Screen   *ebiten.Image
}

func (c *Camera) Update() {
	distance := c.TargetY - c.Position.Y
	if math.Abs(distance) > c.Margin {
		c.Position.Y += distance / c.SpeedReduction
	}
}

func (c *Camera) MoveY(newPos float64) {
	c.TargetY = newPos
}

func (c *Camera) Draw(d DrawOptions) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		60-c.Position.X+d.Position.X-(d.Size.X/2),
		70-c.Position.Y+d.Position.Y-(d.Size.X/2),
	)
	d.Screen.DrawImage(d.Image, op)
}
