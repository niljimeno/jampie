package platforms

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/niljimeno/jampie/utils/vector"
	"github.com/niljimeno/jampie/world"
	"github.com/niljimeno/jampie/world/camera"
)

type Platform struct {
	Image    *ebiten.Image
	Position vector.Vector
	Size     vector.Vector
}

func (p *Platform) Draw(screen *ebiten.Image) {
	world.Camera.Draw(camera.DrawOptions{
		Screen:   screen,
		Image:    p.Image,
		Size:     p.Size,
		Position: p.Position,
	})
}

func NewPlatform() Platform {
	img, _, _ := ebitenutil.NewImageFromFile("assets/platform.png")
	return Platform{
		Image:    img,
		Position: vector.Vector{X: 0, Y: 60},
		Size:     vector.Vector{X: 16, Y: 16},
	}
}
