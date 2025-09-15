package main

import (
	"github.com/niljimeno/jampie/player"
	"github.com/niljimeno/jampie/world"
	"github.com/niljimeno/jampie/world/platforms"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	bat  player.Player
	plat platforms.Platform
)

type Game struct {
}

func (g *Game) init() {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 120, 120
}

func (g *Game) Update() error {
	bat.Update()
	world.Camera.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	bat.Draw(screen)
	plat.Draw(screen)
}

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Attempt")

	world.NewWorld()

	bat, _ = player.NewPlayer()
	plat = platforms.NewPlatform()

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
