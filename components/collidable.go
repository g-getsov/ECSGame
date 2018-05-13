package components

const collidableComponentName = "Collidable"

type Collidable struct {}

func (c Collidable) GetComponentName() string { return GetCollidableComponentName() }

func GetCollidableComponentName() string {return collidableComponentName }