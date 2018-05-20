package components

const healthComponentName = "Health"

type Health struct {
	Health int
	MaxHealth int
}

func (h Health) GetComponentName() string { return GetHealthComponentName() }

func (h Health) IsUniquePerEntity() bool { return true }

func GetHealthComponentName() string {return healthComponentName }