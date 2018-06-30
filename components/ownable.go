package components

const ownableComponentName = "Ownable"

type Ownable struct {
	Id string
	OwnerEntityId string
}

func (o Ownable) GetComponentId() string { return o.Id }

func (o Ownable) GetComponentName() string { return GetOwnableComponentName() }

func (o Ownable) IsUniquePerEntity() bool { return true }

func GetOwnableComponentName() string {return ownableComponentName }