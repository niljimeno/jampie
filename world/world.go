package world

import (
	"github.com/niljimeno/jampie/world/camera"
)

type WorldStruct struct {
	Camera *camera.Camera
}

var Camera camera.Camera

func NewWorld() {
	Camera = camera.Camera{
		SpeedReduction: 8,
		Margin:         1,
	}
}
