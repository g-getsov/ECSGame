package components

const healthComponentName = "Health"

type Health struct {
	Health int
}

func (h Health) GetComponentName() string { return healthComponentName }