package components

func CreateSpeedComponent(speed int) Component {
	return Speed{
		Speed: speed,
	}
}

func CreateVelocityComponent(x int, y int) Component {
	return Velocity{
		X: x,
		Y: y,
	}
}

func CreatePositionComponent(x int, y int) Component {
	return Position{
		X: x,
		Y: y,
	}
}

func CreateHealthComponent(health int) Component {
	return Health{
		Health: health,
	}
}

func CreateInputComponent() Component {
	return Input{}
}