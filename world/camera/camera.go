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

func (c *Camera) Update() {
	distance := c.TargetY - c.Position.Y
	if math.Abs(distance) > c.Margin {
		c.Position.Y += distance / c.SpeedReduction
	}
}

func (c *Camera) SetTargetY(newPos float64) {
	c.TargetY = newPos
}
