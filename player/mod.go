package player

import (
	"github.com/niljimeno/jamping/utils/controls"
	"github.com/niljimeno/jamping/utils/vector"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Image    *ebiten.Image
	Position vector.Vector
	Velocity vector.Vector
	CanJump  bool
}

func (p *Player) Update() {
	var movement float64 = 0
	jumpPressed := false
	if controls.IsKeyPressed("right") {
		movement += 1
	}
	if controls.IsKeyPressed("left") {
		movement -= 1
	}
	if controls.IsKeyPressed("jump") {
		jumpPressed = true
	}

	_ = jumpPressed

	p.Velocity.X += movement
	p.Position.X += p.Velocity.X
	p.Velocity.Y -= 0

	p.Position = vector.Add(p.Position, p.Velocity)
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Position.X, p.Position.Y)
	screen.DrawImage(p.Image, op)
}

func NewPlayer() (Player, error) {
	img, _, err := ebitenutil.NewImageFromFile("assets/bat.png")
	if err != nil {
		return Player{}, err
	}

	return Player{
		Image:    img,
		Position: vector.Vector{X: 0, Y: 0},
	}, nil
}
