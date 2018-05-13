package factories

import (
	rsc "BasicECS/resources"
	"BasicECS/components/weapons"
	"BasicECS/enum"
	"BasicECS/core"
)

type EntityFactory struct {
	imageManager *rsc.ImageManager
}

func CreateEntityFactory(imageManager *rsc.ImageManager) EntityFactory {
	return EntityFactory{
		imageManager,
	}
}

func (e *EntityFactory) CreateWeapon(entityManager *core.EntityManager, x int, y int, fireRate float32, projectileSpeed int, damage int) {

	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("gun.png")

	weaponBaseComponent := CreateWeaponBaseComponent(fireRate, projectileSpeed, damage)
	ownableComponent := CreateOwnableComponent()
	equipableComponent := CreateEquipableComponent(enum.MainHand)
	positionComponent := CreatePositionComponent(x, y)
	spriteComponent := CreateSpriteComponent(image)

	entityManager.AddComponentsToEntity(entity.Id, []core.Component{
		&positionComponent,
		&weaponBaseComponent,
		&ownableComponent,
		&equipableComponent,
		&spriteComponent,
	})
}

func (e *EntityFactory) CreateMagazine(entityManager *core.EntityManager, x int, y int, maxCapacity int) {

	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("magazine.png")

	magazineComponent := CreateMagazineComponent(maxCapacity, nil)
	ownableComponent := CreateOwnableComponent()
	attachableComponent := CreateAttachableComponent(enum.Mag)
	positionComponent := CreatePositionComponent(x, y)
	spriteComponent := CreateSpriteComponent(image)

	entityManager.AddComponentsToEntity(entity.Id, []core.Component{
		&positionComponent,
		&magazineComponent,
		&ownableComponent,
		&attachableComponent,
		&spriteComponent,
	})
}

func (e *EntityFactory) CreateBullet(entityManager *core.EntityManager, x int, y int, speed int, ammunition *weapons.Ammunition) {

	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("player.png")

	positionComponent := CreatePositionComponent(x,y)
	speedComponent := CreateSpeedComponent(speed)
	spriteComponent := CreateSpriteComponent(image)
	movableComponent := CreateMovableComponent()


	entityManager.AddComponentsToEntity(entity.Id, []core.Component{
		&positionComponent,
		&speedComponent,
		&movableComponent,
		&spriteComponent,
	})
}

func (e *EntityFactory) CreatePlayer(entityManager *core.EntityManager, x int, y int, speed int, health int) {
	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("player.png")

	positionComponent := CreatePositionComponent(x,y)
	speedComponent := CreateSpeedComponent(speed)
	healthComponent := CreateHealthComponent(health)
	inputComponent := CreateInputComponent()
	spriteComponent := CreateSpriteComponent(image)
	movableComponent := CreateMovableComponent()

	entityManager.AddComponentsToEntity(entity.Id, []core.Component{
		&positionComponent,
		&speedComponent,
		&movableComponent,
		&healthComponent,
		&inputComponent,
		&spriteComponent,
	})
}

func (e *EntityFactory) CreateZombie(entityManager *core.EntityManager) {
	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("wolf.png")

	positionComponent := CreatePositionComponent(50,50)
	speedComponent := CreateSpeedComponent(1)
	healthComponent := CreateHealthComponent(20)
	spriteComponent := CreateSpriteComponent(image)
	movableComponent := CreateMovableComponent()

	entityManager.AddComponentsToEntity(entity.Id, []core.Component {
		&positionComponent,
		&speedComponent,
		&healthComponent,
		&spriteComponent,
		&movableComponent,
	})
}