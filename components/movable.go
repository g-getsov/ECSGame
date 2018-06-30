package components

const movableComponentName = "Movable"

type Movable struct {
	Id string
	Movable bool
}

func (m Movable) GetComponentId() string { return m.Id }

func (m Movable) GetComponentName() string { return GetMovableComponentName() }

func (m Movable) IsUniquePerEntity() bool { return true }

func GetMovableComponentName() string {return movableComponentName }