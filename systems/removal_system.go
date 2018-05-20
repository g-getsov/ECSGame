package systems

import (
	"BasicECS/core"
	"BasicECS/components"
)

type removalSystem struct {}

func (s removalSystem) Update(dt float64, entityManager *core.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetRemoveComponentName())

	for _, entityId := range entityIds {
		entityManager.RemoveEntity(entityId)
	}
}

func CreateRemovalSystem() System { return removalSystem{} }