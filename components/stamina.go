package components

const staminaComponentName = "Stamina"

type Stamina struct {
	Id string
	Stamina int
	MaxStamina int
}

func (s Stamina) GetComponentId() string { return s.Id }

func (s Stamina) GetComponentName() string { return GetStaminaComponentName() }

func (s Stamina) IsUniquePerEntity() bool { return true }

func GetStaminaComponentName() string {return staminaComponentName }