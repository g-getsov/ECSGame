package weapons

const magazineComponentName = "Magazine"

type Magazine struct {
	Id string
	MaxCapacity int
	Capacity int
	Ammunition *Ammunition
}

func (m Magazine) GetComponentId() string { return m.Id }

func (m Magazine) GetComponentName() string { return GetMagazineComponentName() }

func (m Magazine) IsUniquePerEntity() bool { return true }

func GetMagazineComponentName() string {return magazineComponentName }