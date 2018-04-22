package components

const movableComponentName = "Movable"

type Movable struct {
	Movable bool
}

func (p Movable) GetComponentName() string { return GetMovableComponentName() }

func GetMovableComponentName() string {return movableComponentName }