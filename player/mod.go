package player

import (
	"math"

	"github.com/niljimeno/jampie/utils/controls"
	"github.com/niljimeno/jampie/utils/vector"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Image           *ebiten.Image
	Position        vector.Vector
	Velocity        vector.Vector
	CanJump         bool
	JumpJustPressed bool
	IsOnGround      bool
}

func (p *Player) Update() {
	jumpPressed := false

	var movement int = 0
	var speed float64 = 0.05
	var gravity float64 = 0.2

	if controls.IsKeyPressed("right") {
		movement += 1
	}
	if controls.IsKeyPressed("left") {
		movement -= 1
	}

	if keypressed := controls.IsKeyPressed("jump"); keypressed && p.JumpJustPressed == false {
		jumpPressed = true
		p.JumpJustPressed = true
	} else {
		p.JumpJustPressed = keypressed
	}

	if movement != 0 {
		p.Velocity.X += speed * float64(movement)
	}

	if p.IsOnGround {
		if (movement == 0) || ((movement > 0) != (p.Velocity.X > 0)) {
			switch {
			case p.Velocity.X < 0:
				p.Velocity.X += speed
			case p.Velocity.X > 0:
				p.Velocity.X -= speed
			}
		}

	} else {
		if p.Position.Y <= 69 {
			p.Velocity.Y += gravity
		} else {
			p.Velocity.Y = 0
			p.IsOnGround = true
		}
	}

	if movement == 0 && math.Abs(p.Velocity.X) < 0.2 {
		p.Velocity.X = 0
	}

	if jumpPressed {
		p.Velocity.Y = -2.69
		p.IsOnGround = false
	}

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
		Image:      img,
		Position:   vector.Vector{X: 0, Y: 0},
		IsOnGround: false,
	}, nil
}
