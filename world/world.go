package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/niljimeno/jampie/world/camera"
	"github.com/niljimeno/jampie/world/terrain"
)

type WorldStruct struct {
	Camera *camera.Camera
}

var Camera camera.Camera
var Terrain terrain.Terrain

func NewWorld() {
	Camera = camera.Camera{
		SpeedReduction: 16,
		Margin:         1,
	}
}

func GenerateTerrain() error {
	var err error
	Terrain, err = terrain.GenerateTerrain("assets/maps/template.png")
	if err != nil {
		return err
	}

	return nil
}

func DrawWorld(screen *ebiten.Image) {
	Terrain.Draw(screen, Camera)
}
