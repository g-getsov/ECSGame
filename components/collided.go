package components

const collidedComponentName = "Collided"

type Collided struct {
	CollidedEntities map[string]bool
}

func (c Collided) GetComponentName() string { return GetCollidedComponentName() }

func (c Collided) IsUniquePerEntity() bool { return true }

func GetCollidedComponentName() string {return collidedComponentName }
