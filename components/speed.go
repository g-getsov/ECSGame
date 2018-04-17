package components

const speedComponentName = "Speed"

type Speed struct {
	Speed int
}

func (s Speed) GetComponentName() string { return GetSpeedComponentName() }

func GetSpeedComponentName() string {return speedComponentName }