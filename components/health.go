package components

const healthComponentName = "Health"

type Health struct {
	Health int
	MaxHealth int
}

func (h Health) GetComponentName() string { return GetHealthComponentName() }

func GetHealthComponentName() string {return healthComponentName }