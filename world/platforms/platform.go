package platforms

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/niljimeno/jampie/utils/vector"
	"github.com/niljimeno/jampie/world"
)

type Platform struct {
	Position vector.Vector
	Size     vector.Vector
}

func (p *Platform) Draw(screen *ebiten.Image) {
	world.Camera.Pix(screen, p.Position, color.White)
}

func NewPlatform() Platform {
	return Platform{
		Position: vector.Vector{X: 20, Y: 60},
		Size:     vector.Vector{X: 16, Y: 16},
	}
}
