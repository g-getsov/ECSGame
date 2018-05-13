package components

const expirableComponentName = "Expirable"

type Expirable struct {
	TimeLeft int
}

func (e Expirable) GetComponentName() string { return GetExpirableComponentName() }

func GetExpirableComponentName() string { return expirableComponentName }
