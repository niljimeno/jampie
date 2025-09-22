package terrain

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/niljimeno/jampie/settings"
	"github.com/niljimeno/jampie/utils/vector"
	"github.com/niljimeno/jampie/world/camera"
)

const terrainWidth = settings.Width

type Terrain [][terrainWidth]byte

var image *ebiten.Image

type rgba32 struct {
	R, G, B, A uint32
}

var (
	grass = rgba32{
		R: 0x6767,
		G: 0xaaaa,
		B: 0x2828,
		A: 0xffff,
	}

	ground = rgba32{
		R: 0x3636,
		G: 0x3131,
		B: 0x1616,
		A: 0xffff,
	}
)

func (t *Terrain) Draw(screen *ebiten.Image, cam camera.Camera) {
	cam.Draw(camera.DrawOptions{
		Position: vector.Vector{X: 0, Y: 0},
		Size:     vector.Vector{X: settings.Width, Y: settings.Height},
		Image:    image,
		Screen:   screen,
	})
}

func GenerateTerrain(source string) (Terrain, error) {
	var err error
	image, _, err = ebitenutil.NewImageFromFile(source)
	if err != nil {
		return nil, err
	}

	terrain := [][terrainWidth]byte{}

	for r := range image.Bounds().Max.Y {
		terrain = append(terrain, [128]byte{})
		for i := range 128 {
			x := i
			y := r

			imageColor := intoRGB32(image.At(x, y))

			switch {
			case compareColor(grass, imageColor):
				terrain[r][i] = 1
			case compareColor(ground, imageColor):
				terrain[r][i] = 2
			default:
				terrain[r][i] = 0
			}
		}
	}

	return terrain, nil
}

func intoRGB32(input color.Color) rgba32 {
	r, g, b, a := input.RGBA()
	return rgba32{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func compareColor(a, b rgba32) bool {
	return (a.R == b.R &&
		a.G == b.G &&
		a.B == b.B &&
		a.A == b.A)
}
