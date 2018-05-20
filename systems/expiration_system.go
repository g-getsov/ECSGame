package systems

import (
	"BasicECS/core"
	"BasicECS/components"
	"BasicECS/factories"
)

type expirationSystem struct {}

func (s expirationSystem) Update(dt float64, entityManager *core.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(components.GetExpirableComponentName())

	for _, entityId := range entityIds {

		expirableComponent, _ := entityManager.GetComponentOfClass(
			components.GetExpirableComponentName(),
			entityId).(*components.Expirable)

		ownableComponent, _ := entityManager.GetComponentOfClass(
			components.GetOwnableComponentName(),
			entityId).(*components.Ownable)

		//don't expire while owned
		if ownableComponent != nil {
			if len(ownableComponent.OwnerEntityId) == 0 {
				expirableComponent.TimeLeft -= dt
			}
		} else {
			expirableComponent.TimeLeft -= dt
		}

		if expirableComponent.TimeLeft <= 0 {
			entityManager.AddComponentToEntity(
				entityId,
				factories.CreateRemoveComponent())

			entityManager.RemoveComponentForEntity(entityId, components.GetExpirableComponentName())
		}
	}
}

func CreateExpirationSystem() System { return expirationSystem{} }