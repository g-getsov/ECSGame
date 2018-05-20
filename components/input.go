package components

const inputComponentName = "Input"

type Input struct {
	ForwardKey bool
	BackwardsKey bool
	LeftKey bool
	RightKey bool
	ShootKey bool
	InteractKey bool
	InventoryToggle bool
}

func (i Input) GetComponentName() string { return GetInputComponentName() }

func GetInputComponentName() string {return inputComponentName }