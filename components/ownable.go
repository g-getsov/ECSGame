package components

const ownableComponentName = "Ownable"

type Ownable struct {
	OwnerEntityId string
}

func (o Ownable) GetComponentName() string { return GetOwnableComponentName() }

func GetOwnableComponentName() string {return ownableComponentName }