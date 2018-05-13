package entities

import (
	cmpt "BasicECS/components"
	rsc "BasicECS/resources"
	"BasicECS/components/weapons"
	"BasicECS/enum"
)

type EntityFactory struct {
	imageManager *rsc.ImageManager
}

func CreateEntityFactory(imageManager *rsc.ImageManager) EntityFactory {
	return EntityFactory{
		imageManager,
	}
}

func (e *EntityFactory) CreateWeapon(entityManager *EntityManager, x int, y int, fireRate float32, projectileSpeed int, damage int) {

	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("gun.png")

	weaponBaseComponent := cmpt.CreateWeaponBaseComponent(fireRate, projectileSpeed, damage)
	ownableComponent := cmpt.CreateOwnableComponent()
	equipableComponent := cmpt.CreateEquipableComponent(enum.MainHand)
	positionComponent := cmpt.CreatePositionComponent(x, y)
	spriteComponent := cmpt.CreateSpriteComponent(image)

	entityManager.AddComponentsToEntity(entity.Id, []cmpt.Component{
		&positionComponent,
		&weaponBaseComponent,
		&ownableComponent,
		&equipableComponent,
		&spriteComponent,
	})
}

func (e *EntityFactory) CreateMagazine(entityManager *EntityManager, x int, y int, maxCapacity int) {

	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("magazine.png")

	magazineComponent := cmpt.CreateMagazineComponent(maxCapacity, nil)
	ownableComponent := cmpt.CreateOwnableComponent()
	attachableComponent := cmpt.CreateAttachableComponent(enum.Mag)
	positionComponent := cmpt.CreatePositionComponent(x, y)
	spriteComponent := cmpt.CreateSpriteComponent(image)

	entityManager.AddComponentsToEntity(entity.Id, []cmpt.Component{
		&positionComponent,
		&magazineComponent,
		&ownableComponent,
		&attachableComponent,
		&spriteComponent,
	})
}

func (e *EntityFactory) CreateBullet(entityManager *EntityManager, x int, y int, speed int, ammunition *weapons.Ammunition) {

	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("player.png")

	positionComponent := cmpt.CreatePositionComponent(x,y)
	speedComponent := cmpt.CreateSpeedComponent(speed)
	spriteComponent := cmpt.CreateSpriteComponent(image)
	movableComponent := cmpt.CreateMovableComponent()


	entityManager.AddComponentsToEntity(entity.Id, []cmpt.Component{
		&positionComponent,
		&speedComponent,
		&movableComponent,
		&spriteComponent,
	})
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