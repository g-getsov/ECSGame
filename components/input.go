package components

const inputComponentName = "Input"

type Input struct {
	Id string
	ForwardKey bool
	BackwardsKey bool
	LeftKey bool
	RightKey bool
	ShootKey bool
	InteractKey bool
	InventoryToggle bool
}

func (i Input) GetComponentId() string { return i.Id }

func (i Input) GetComponentName() string { return GetInputComponentName() }

func (i Input) IsUniquePerEntity() bool { return true }

func GetInputComponentName() string {return inputComponentName }