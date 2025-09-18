package terrain

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/niljimeno/jampie/world/camera"
)

type Terrain [][128]byte

type rgba32 struct {
	R, G, B, A uint32
}

var (
	ground = rgba32{
		R: 000,
		G: 000,
		B: 000,
		A: 255,
	}

	air = rgba32{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
)

func (t *Terrain) Draw(screen *ebiten.Image, cam camera.Camera) {
	for y, line := range *t {
		for x, pix := range line {
			var col color.Color
			col = color.White
			switch pix {
			case 0:
				col = color.Black
			case 1:
				col = color.White
			}

			cam.Pix(
				screen,
				x-60,
				-y,
				col,
			)
		}
	}
}

func GenerateTerrain(source string) (Terrain, error) {
	img, _, err := ebitenutil.NewImageFromFile(source)
	if err != nil {
		return nil, err
	}

	terrain := [][128]byte{{}}
	fmt.Println("les go", img.Bounds().Max.X)

	for r := range img.Bounds().Max.Y {
		for i := range img.Bounds().Max.X {
			x := i
			y := r

			imageColor := intoRGB32(img.At(x, y))

			fmt.Println(imageColor)

			switch {
			default:
				terrain[r][i] = 0
			case compareColor(air, imageColor):
				terrain[r][i] = 0
			case compareColor(ground, imageColor):
				terrain[r][i] = 1
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
