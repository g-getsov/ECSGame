package components

const healthComponentName = "Health"

type Health struct {
	Id string
	Health int
	MaxHealth int
}

func (h Health) GetComponentId() string { return h.Id }

func (h Health) GetComponentName() string { return GetHealthComponentName() }

func (h Health) IsUniquePerEntity() bool { return true }

func GetHealthComponentName() string {return healthComponentName }