package components

import "github.com/hajimehoshi/ebiten"

func CreateSpeedComponent(speed int) Speed {
	return Speed{
		Speed: speed,
	}
}

func CreateVelocityComponent(x int, y int) Velocity {
	return Velocity{
		X: x,
		Y: y,
	}
}

func CreatePositionComponent(x int, y int) Position {
	return Position{
		X: x,
		Y: y,
	}
}

func CreateHealthComponent(health int) Health {
	return Health{
		Health: health,
	}
}

func CreateInputComponent() Input {
	return Input{}
}

func CreateSpriteComponent(image *ebiten.Image) Sprite {
	return Sprite{
		Image: image,
	}
}

func CreateMovableComponent() Movable {
	return Movable{ true }
}