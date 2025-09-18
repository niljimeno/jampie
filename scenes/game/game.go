package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/niljimeno/jampie/player"
	"github.com/niljimeno/jampie/world"
)

var bat player.Player

func Init() {
	bat, _ = player.NewPlayer()
	err := world.GenerateTerrain()
	if err != nil {
		log.Println(err)
	}
}

func Update(nextScene *int) {
	bat.Update()
	world.Camera.Update()
}

func Draw(screen *ebiten.Image) {
	bat.Draw(screen)
	world.DrawWorld(screen)
}
