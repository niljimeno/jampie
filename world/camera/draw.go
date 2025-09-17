package camera

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	ebitenVector "github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/niljimeno/jampie/utils/vector"
)

const (
	centerX = 60
	centerY = 60
)

type DrawOptions struct {
	Position vector.Vector
	Size     vector.Vector
	Image    *ebiten.Image
	Screen   *ebiten.Image
}

func (c *Camera) Draw(d DrawOptions) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		centerX-c.Position.X+d.Position.X-(d.Size.X/2),
		centerY-c.Position.Y+d.Position.Y-(d.Size.X/2),
	)
	d.Screen.DrawImage(d.Image, op)
}

func (c *Camera) Pix(screen *ebiten.Image, position vector.Vector, col color.Color) {
	ebitenVector.DrawFilledRect(
		screen,
		float32(position.X-c.Position.X+centerX),
		float32(position.Y-c.Position.Y+centerY),
		1, 1,
		col,
		false,
	)
}
