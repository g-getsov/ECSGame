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
	Type DamageType
	Damage int
}

func (v Damage) GetComponentName() string { return GetDamageComponentName() }

func GetDamageComponentName() string {return damageComponentName }