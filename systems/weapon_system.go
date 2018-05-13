package systems

/*import (
	"github.com/hajimehoshi/ebiten"
	"BasicECS/entities"
	"BasicECS/components"
)

type weaponSystem struct {
	screen *ebiten.Image
}

func (s weaponSystem) Update(dt float64, entityManager *entities.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetWeaponComponentName())

	for _, entityId := range entityIds {

		weaponComponent := entityManager.GetComponentOfClass(
			components.GetWeaponComponentName(),
			entityId).(*components.Weapon)

		if weaponComponent.Fired {
			magazine := weaponComponent.Magazine
			if magazine.Capacity > 0 {
				magazine.Capacity -= 1

			} else {
				if weaponComponent.NumMags <= 0 { continue }
				weaponComponent.NumMags -= 1
				magazine.Capacity = magazine.MaxCapacity
			}
		}
	}
}

func CreateWeaponSystem() System { return weaponSystem{} }*/