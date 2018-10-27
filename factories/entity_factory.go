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

func (e *EntityFactory) CreateWeapon(entityManager *core.EntityManager, x int, y int, fireRate float32, projectileSpeed int, damage int) string {

	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("gun.png")
	thumbnail := e.imageManager.GetImageOrLoad("gun_thumbnail.png")

	weaponBaseComponent := CreateWeaponBaseComponent(fireRate, projectileSpeed, damage)
	ownableComponent := CreateOwnableComponent()
	expirableComponent := CreateExpirableComponent(10)
	equipableComponent := CreateEquipableComponent(enum.MainHand)
	positionComponent := CreatePositionComponent(x, y)
	spriteComponent := CreateSpriteComponent(image, thumbnail)
	tooltipComponent := CreateTooltipComponent("Press 'E' to pick up")
	hitboxComponent := CreateHitboxComponent(64,64, false)
	interactiveComponent := CreateInteractiveComponent(PickUp)

	entityManager.AddComponentsToEntity(entity.Id, []core.Component{
		&positionComponent,
		&weaponBaseComponent,
		&ownableComponent,
		&equipableComponent,
		&spriteComponent,
		&expirableComponent,
		&tooltipComponent,
		&hitboxComponent,
		&interactiveComponent,
	})

	return entity.Id
}

func (e *EntityFactory) CreateMagazine(entityManager *core.EntityManager, x int, y int, maxCapacity int) string {

	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("magazine.png")
	thumbnail := e.imageManager.GetImageOrLoad("missing_thumbnail.png")

	magazineComponent := CreateMagazineComponent(maxCapacity, nil)
	ownableComponent := CreateOwnableComponent()
	attachableComponent := CreateAttachableComponent(enum.Mag)
	positionComponent := CreatePositionComponent(x, y)
	spriteComponent := CreateSpriteComponent(image, thumbnail)
	tooltipComponent := CreateTooltipComponent("Press 'E' to pick up")
	hitboxComponent := CreateHitboxComponent(64,64, false)
	interactiveComponent := CreateInteractiveComponent(PickUp)

	entityManager.AddComponentsToEntity(entity.Id, []core.Component{
		&positionComponent,
		&magazineComponent,
		&ownableComponent,
		&attachableComponent,
		&spriteComponent,
		&tooltipComponent,
		&hitboxComponent,
		&interactiveComponent,
	})

	return entity.Id
}

func (e *EntityFactory) CreateBullet(entityManager *core.EntityManager, x int, y int, speed int, ammunition *weapons.Ammunition) string {

	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("player.png")

	positionComponent := CreatePositionComponent(x,y)
	speedComponent := CreateSpeedComponent(speed)
	spriteComponent := CreateSpriteComponent(image, nil)
	movableComponent := CreateMovableComponent()


	entityManager.AddComponentsToEntity(entity.Id, []core.Component{
		&positionComponent,
		&speedComponent,
		&movableComponent,
		&spriteComponent,
	})

	return entity.Id
}

func (e *EntityFactory) CreatePlayer(entityManager *core.EntityManager, x int, y int, speed int, health int, stamina int) string {
	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("player.png")

	positionComponent := CreatePositionComponent(x,y)
	speedComponent := CreateSpeedComponent(speed)
	healthComponent := CreateHealthComponent(health)
	staminaComponent := CreateStaminaComponent(stamina)
	inputComponent := CreateInputComponent()
	spriteComponent := CreateSpriteComponent(image, nil)
	movableComponent := CreateMovableComponent()
	hitboxComponent := CreateHitboxComponent(64, 64, true)
	inventoryComponent := CreateInventoryComponent(100)

	entityManager.AddComponentsToEntity(entity.Id, []core.Component{
		&positionComponent,
		&speedComponent,
		&movableComponent,
		&healthComponent,
		&staminaComponent,
		&inputComponent,
		&spriteComponent,
		&hitboxComponent,
		&inventoryComponent,
	})

	return entity.Id
}

func (e *EntityFactory) CreateZombie(entityManager *core.EntityManager) string {
	entity := entityManager.CreateEntity()

	image := e.imageManager.GetImageOrLoad("wolf.png")

	positionComponent := CreatePositionComponent(50,50)
	speedComponent := CreateSpeedComponent(1)
	healthComponent := CreateHealthComponent(20)
	spriteComponent := CreateSpriteComponent(image, nil)
	movableComponent := CreateMovableComponent()
	hitboxComponent := CreateHitboxComponent(64, 64, true)
	tooltipComponent := CreateTooltipComponent("Grrrr... !")

	entityManager.AddComponentsToEntity(entity.Id, []core.Component {
		&positionComponent,
		&speedComponent,
		&healthComponent,
		&spriteComponent,
		&movableComponent,
		&hitboxComponent,
		&tooltipComponent,
	})

	return entity.Id
}

func (e *EntityFactory) CreateGui(entityManager *core.EntityManager, entityId string) string {

	entity := entityManager.CreateEntity()

	guiComponent := CreateGuiComponent(entityId)
	characterStatsGuiComponent := CreateCharacterStatsGuiComponent()
	inventoryGuiComponent := CreateInventoryGuiComponent()

	entityManager.AddComponentsToEntity(entity.Id, []core.Component {
		&guiComponent,
		&characterStatsGuiComponent,
		&inventoryGuiComponent,
	})

	return entity.Id
}