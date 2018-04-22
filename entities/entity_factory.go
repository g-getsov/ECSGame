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

func (e *EntityFactory) CreatePlayer(entityManager *EntityManager) {
	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("player.png")

	positionComponent := cmpt.CreatePositionComponent(200,200)
	speedComponent := cmpt.CreateSpeedComponent(2)
	healthComponent := cmpt.CreateHealthComponent(100)
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

	positionComponent := cmpt.CreatePositionComponent(250,250)
	speedComponent := cmpt.CreateSpeedComponent(1)
	healthComponent := cmpt.CreateHealthComponent(20)

	entityManager.AddComponentsToEntity(entity.Id, []cmpt.Component {
		&positionComponent,
		&speedComponent,
		&healthComponent,
	})
}