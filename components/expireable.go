package components

const expirableComponentName = "Expirable"

type Expirable struct {
	Id string
	TimeLeft float64
	componentName string
}

func (e Expirable) GetComponentId() string { return e.Id }

func (e Expirable) GetComponentName() string { return GetExpirableComponentName() }

func (e Expirable) IsUniquePerEntity() bool { return true }

func GetExpirableComponentName() string { return expirableComponentName }
