package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/niljimeno/jampie/scenes"
)

func Update(nextScene *int) {
	*nextScene = scenes.Game
}

func Draw(screen *ebiten.Image) {
}
