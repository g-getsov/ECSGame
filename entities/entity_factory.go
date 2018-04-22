package entities

import (
	cmpt "BasicECS/components"
	rsc "BasicECS/resources"
)

type EntityFactory struct {
	imageManager *rsc.ImageManager
}

func CreateEntityFactory(imageManager *rsc.ImageManager) EntityFactory {
	return EntityFactory{
		imageManager,
	}
}

func (e *EntityFactory) CreatePlayer(entityManager *EntityManager, x int, y int, speed int, health int) {
	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("player.png")

	positionComponent := cmpt.CreatePositionComponent(x,y)
	speedComponent := cmpt.CreateSpeedComponent(speed)
	healthComponent := cmpt.CreateHealthComponent(health)
	inputComponent := cmpt.CreateInputComponent()
	spriteComponent := cmpt.CreateSpriteComponent(image)
	movableComponent := cmpt.CreateMovableComponent()

	entityManager.AddComponentsToEntity(entity.Id, []cmpt.Component{
		&positionComponent,
		&speedComponent,
		&movableComponent,
		&healthComponent,
		&inputComponent,
		&spriteComponent,
	})
}

func (e *EntityFactory) CreateZombie(entityManager *EntityManager) {
	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("wolf.png")

	positionComponent := cmpt.CreatePositionComponent(50,50)
	speedComponent := cmpt.CreateSpeedComponent(1)
	healthComponent := cmpt.CreateHealthComponent(20)
	spriteComponent := cmpt.CreateSpriteComponent(image)
	movableComponent := cmpt.CreateMovableComponent()

	entityManager.AddComponentsToEntity(entity.Id, []cmpt.Component {
		&positionComponent,
		&speedComponent,
		&healthComponent,
		&spriteComponent,
		&movableComponent,
	})
}