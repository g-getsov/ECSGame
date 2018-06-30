package weapons

const weaponBaseName = "WeaponBase"

type WeaponBase struct {
	Id string
	FireRate float32
	ProjectileSpeed int
	Damage int
}

func (w WeaponBase) GetComponentId() string { return w.Id }

func (w WeaponBase) GetComponentName() string { return GetWeaponBaseComponentName() }

func (w WeaponBase) IsUniquePerEntity() bool { return true }

func GetWeaponBaseComponentName() string {return weaponBaseName }