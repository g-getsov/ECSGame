package weapons

const ammunitionComponentName = "Ammunition"

type AmmunitionType int

const (
	NineMilimiter AmmunitionType = 1
	TwelveGauge AmmunitionType = 2
	Energy AmmunitionType = 3
)

type Ammunition struct {
	Id string
	Damage int
	AmmunitionType AmmunitionType
}

func (a Ammunition) GetComponentId() string { return a.Id }

func (a Ammunition) GetComponentName() string { return GetAmmunitionComponentName() }

func (a Ammunition) IsUniquePerEntity() bool { return true }

func GetAmmunitionComponentName() string {return ammunitionComponentName }