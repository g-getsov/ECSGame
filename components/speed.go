package components

const speedComponentName = "Speed"

type Speed struct {
	Id string
	Speed int
}

func (s Speed) GetComponentId() string { return s.Id }

func (s Speed) GetComponentName() string { return GetSpeedComponentName() }

func (s Speed) IsUniquePerEntity() bool { return true }

func GetSpeedComponentName() string {return speedComponentName }