package main

import (
	"github.com/niljimeno/jamping/player"

	"github.com/hajimehoshi/ebiten/v2"
)

var bat player.Player

type Game struct {
}

func (g *Game) init() {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 100, 100
}

func (g *Game) Update() error {
	bat.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	bat.Draw(screen)
}

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Attempt")

	bat, _ = player.NewPlayer()

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
