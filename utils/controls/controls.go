package controls

import "github.com/hajimehoshi/ebiten/v2"

func IsKeyPressed(key string) bool {
	switch key {
	default:
		return false
	case "jump":
		return (ebiten.IsKeyPressed(ebiten.KeyUp) ||
			ebiten.IsKeyPressed(ebiten.KeyK) ||
			ebiten.IsKeyPressed(ebiten.KeyW))
	case "left":
		return (ebiten.IsKeyPressed(ebiten.KeyLeft) ||
			ebiten.IsKeyPressed(ebiten.KeyH) ||
			ebiten.IsKeyPressed(ebiten.KeyA))
	case "right":
		return (ebiten.IsKeyPressed(ebiten.KeyRight) ||
			ebiten.IsKeyPressed(ebiten.KeyL) ||
			ebiten.IsKeyPressed(ebiten.KeyD))
	case "quit":
		return (ebiten.IsKeyPressed(ebiten.KeyQ) ||
			ebiten.IsKeyPressed(ebiten.KeyEscape))
	}
}
