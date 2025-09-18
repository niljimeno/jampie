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
	Scene     int
	NextScene int
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return settings.Width, settings.Height
}

func (g *Game) ChangeScene() {
	g.Scene = g.NextScene

	switch g.Scene {
	case scenes.Menu:
	case scenes.Game:
		game.Init()
	}
}

func (g *Game) Update() error {
	if g.Scene != g.NextScene {
		g.ChangeScene()
	}

	switch g.Scene {
	case scenes.Menu:
		menu.Update(&g.NextScene)
	case scenes.Game:
		game.Update(&g.NextScene)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.Scene {
	case scenes.Menu:
		menu.Draw(screen)
	case scenes.Game:
		game.Draw(screen)
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

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
