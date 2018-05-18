package systems

import (
	cmpt "BasicECS/components"
	"BasicECS/core"
	"strings"
	"BasicECS/factories"
)

type collisionSystem struct {}

func (s collisionSystem) Update(dt float64, entityManager *core.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(cmpt.GetHitboxComponentName())

	for _, entityId := range entityIds {

		hitboxComponent, _ := entityManager.GetComponentOfClass(
			cmpt.GetHitboxComponentName(),
			entityId).(*cmpt.Hitbox)

		positionComponent, _ := entityManager.GetComponentOfClass(
			cmpt.GetPositionComponentName(),
			entityId).(*cmpt.Position)

		if positionComponent == nil { continue }

		entityManager.RemoveComponentForEntity(entityId, cmpt.GetCollidedComponentName())
		//hasCollidedWithAnything := false

		for _, collidingEntityId := range entityIds {

			if strings.Compare(entityId, collidingEntityId) == 0 { continue }

			collidingHitboxComponent, _ := entityManager.GetComponentOfClass(
				cmpt.GetHitboxComponentName(),
				collidingEntityId).(*cmpt.Hitbox)

			collidingPositionComponent, _ := entityManager.GetComponentOfClass(
				cmpt.GetPositionComponentName(),
				collidingEntityId).(*cmpt.Position)

			if collidingPositionComponent == nil { continue }

			doHitboxsesCollide :=
				doHitboxesColideAABB(
					positionComponent,
					hitboxComponent,
					collidingPositionComponent,
					collidingHitboxComponent)

			if !doHitboxsesCollide { continue }

			//hasCollidedWithAnything = true

			addCollision(entityManager, entityId, collidingEntityId) //add collision for first colliding object
/*			addCollision(entityManager, collidingEntityId, entityId)*/
		}

		/*if !hasCollidedWithAnything {
			entityManager.RemoveComponentForEntity(entityId, cmpt.GetCollidedComponentName())
		}*/
	}
}

func CreateCollisionSystem() System { return collisionSystem{} }

func doHitboxesColideAABB(positionOne *cmpt.Position, hitboxOne *cmpt.Hitbox, positionTwo *cmpt.Position, hitboxTwo *cmpt.Hitbox) bool {
	return positionOne.X < positionTwo.X + hitboxTwo.Width &&
		positionTwo.X < positionOne.X + hitboxOne.Width &&
		positionOne.Y <	positionTwo.Y + hitboxTwo.Height &&
		positionTwo.Y < positionOne.Y + hitboxOne.Height
}

func addCollision(entityManager *core.EntityManager, entityId string, collidingEntityId string) {

	colliedComponent, _ := entityManager.GetComponentOfClass(
		cmpt.GetCollidedComponentName(),
		entityId).(*cmpt.Collided)

	if colliedComponent == nil {
		newCollidedComponent := factories.CreateCollidedComponent()
		colliedComponent = &newCollidedComponent
	} else {
		if	_, exists := colliedComponent.CollidedEntities[collidingEntityId]; exists { return }
	}

	colliedComponent.CollidedEntities[collidingEntityId] = true

	entityManager.AddComponentToEntity(entityId, colliedComponent)
}