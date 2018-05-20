package components

const expirableComponentName = "Expirable"

type Expirable struct {
	TimeLeft float64
	componentName string
}

func (e Expirable) GetComponentName() string { return GetExpirableComponentName() }

func GetExpirableComponentName() string { return expirableComponentName }
