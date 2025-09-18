package main

import (
	"github.com/niljimeno/jampie/scenes"
	"github.com/niljimeno/jampie/scenes/game"
	"github.com/niljimeno/jampie/scenes/menu"
	"github.com/niljimeno/jampie/settings"
	"github.com/niljimeno/jampie/world"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Scene int
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return settings.Width, settings.Height
}

func (g *Game) Update() error {
	switch g.Scene {
	case scenes.Menu:
		menu.Update()
	case scenes.Game:
		game.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.Scene {
	case scenes.Menu:
		menu.Update()
	case scenes.Game:
		game.Update()
	}
}

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Attempt")

	world.NewWorld()

	runGame()
}

func runGame() {
	g := &Game{}
	g.Scene = scenes.Menu
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
