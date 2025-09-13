package player

import (
	"math"

	"github.com/niljimeno/jampie/utils/controls"
	"github.com/niljimeno/jampie/utils/imath"
	"github.com/niljimeno/jampie/utils/vector"
	"github.com/niljimeno/jampie/world"
	"github.com/niljimeno/jampie/world/camera"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Image      *ebiten.Image
	Position   vector.Vector
	Velocity   vector.Vector
	Size       vector.Vector
	CanJump    bool
	IsOnGround bool

	Ready      bool
	Input      PlayerInput
	Properties PlayerProperties
}

type PlayerInput struct {
	Move            int
	Jump            bool
	JumpJustPressed bool
}

type PlayerProperties struct {
	BASE_GRAVITY float64
	BASE_SPEED   float64
	Gravity      float64
	Speed        float64
}

func (p *Player) Update() {
	if !p.Ready {
		p.SetProperties(0.09, 0.05)
	}

	p.HandleInput()
	p.HandlePhysics()

	p.Position = vector.Add(p.Position, p.Velocity)
}

func (p *Player) HandlePhysics() {
	if p.IsOnGround {
		p.GroundPhysics()
	} else {
		p.AirPhysics()
	}

	switch {
	case p.Input.Move != 0:
		var movement = p.Velocity.X + (p.Properties.Speed * float64(p.Input.Move))
		p.Velocity.X = imath.Clamp(movement, -1.69, 1.69)

	case p.Input.Move == 0 && math.Abs(p.Velocity.X) < 0.25:
		p.Velocity.X = 0
	}

	if p.Input.Jump {
		p.Velocity.Y = -2.0
		p.IsOnGround = false
	}
}

func (p *Player) GroundPhysics() {
	if (p.Input.Move == 0) || ((p.Input.Move > 0) != (p.Velocity.X > 0)) {
		switch {
		case p.Velocity.X < 0:
			p.Velocity.X += p.Properties.Speed * 2
		case p.Velocity.X > 0:
			p.Velocity.X -= p.Properties.Speed * 2
		}
	}
}

func (p *Player) AirPhysics() {
	if p.Position.Y <= 100 {
		p.Velocity.Y += p.Properties.Gravity
	} else {
		p.Velocity.Y = 0
		p.Position.Y = 100
		p.IsOnGround = true
	}
}

func (p *Player) HandleInput() {
	p.Input.Jump = false
	p.Input.Move = 0

	if controls.IsKeyPressed("right") {
		p.Input.Move = 1
	}
	if controls.IsKeyPressed("left") {
		p.Input.Move = -1
	}

	if keypressed := controls.IsKeyPressed("jump"); keypressed && p.Input.JumpJustPressed == false {
		p.Input.Jump = true
		p.Input.JumpJustPressed = true
	} else {
		p.Input.JumpJustPressed = keypressed
	}
}

func (p *Player) SetProperties(gravity float64, speed float64) {
	p.Properties.BASE_GRAVITY = gravity
	p.Properties.BASE_SPEED = speed

	p.Properties.Gravity = gravity
	p.Properties.Speed = speed

	p.Ready = true
}

func (p *Player) Draw(screen *ebiten.Image) {
	world.Camera.Draw(camera.DrawOptions{
		Position: p.Position,
		Size:     p.Size,
		Image:    p.Image,
	})

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
		Ready:      false,
	}, nil
}
