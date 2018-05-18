package components

const expirableComponentName = "Expirable"

type Expirable struct {
	TimeLeft float64
}

func (e Expirable) GetComponentName() string { return GetExpirableComponentName() }

func GetExpirableComponentName() string { return expirableComponentName }
