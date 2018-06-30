package components

const damageComponentName = "Damage"

type DamageType int

const (
	Standard DamageType = 1
	Incendiary DamageType = 2
	Poisonous DamageType = 3
	Electric DamageType = 4
	Magic DamageType = 5
)

type Damage struct {
	Id string
	Type DamageType
	Damage int
}

func (d Damage) GetComponentId() string { return d.Id }

func (d Damage) GetComponentName() string { return GetDamageComponentName() }

func (d Damage) IsUniquePerEntity() bool { return true }

func GetDamageComponentName() string {return damageComponentName }