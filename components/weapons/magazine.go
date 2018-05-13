package weapons

const magazineComponentName = "Magazine"

type Magazine struct {
	MaxCapacity int
	Capacity int
	Ammunition *Ammunition
}

func (v Magazine) GetComponentName() string { return GetMagazineComponentName() }

func GetMagazineComponentName() string {return magazineComponentName }