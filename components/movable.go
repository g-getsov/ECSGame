package components

const movableComponentName = "Movable"

type Movable struct {
	Movable bool
}

func (m Movable) GetComponentName() string { return GetMovableComponentName() }

func (m Movable) IsUniquePerEntity() bool { return true }

func GetMovableComponentName() string {return movableComponentName }