package systems

import (
	"BasicECS/core"
	"BasicECS/components/guis"
	"BasicECS/components"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"BasicECS/enum"
)

type guiRenderSystem struct {
	screen *ebiten.Image
}

const INVENTORY_WINDOW_WIDTH = INVENTORY_MAX_CELLS_IN_ROW*INVENTORY_CELL_WIDTH
const INVENTORY_WINDOW_HEIGHT = 300
const INVENTORY_WINDOW_POSITION_X = 350
const INVENTORY_WINDOW_POSITION_Y = 100

const INVENTORY_MAX_CELLS_IN_ROW = 8
const INVENTORY_CELL_WIDTH = 32
const INVENTORY_CELL_HEIGHT = 32

const HEALTH_BAR_POSITION_X = 30
const HEALTH_BAR_POSITION_Y = 30
const HEALTH_BAR_MAX_WIDTH = 200
const HEALTH_BAR_HEIGHT = 10

func (s guiRenderSystem) Update(dt float64, entityManager *core.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(guis.GetGuiComponentName())

	for _, guiEntityId := range entityIds {

		guiComponent, _ := entityManager.GetComponentOfClass(
		guis.GetGuiComponentName(),
			guiEntityId).(*guis.Gui)

		entityId := guiComponent.EntityId

		characterStatsGuiComponent, _ := entityManager.GetComponentOfClass(
			guis.GetCharacterStatsGuiComponentName(),
			guiEntityId).(*guis.CharacterStatsGui)

		if characterStatsGuiComponent != nil {
			s.renderCharacterStatsGui(entityManager, entityId)
		}

		inventoryGuiComponent, _ := entityManager.GetComponentOfClass(
			guis.GetInvnetoryGuiComponentName(),
			guiEntityId).(*guis.InventoryGui)

		if inventoryGuiComponent != nil {
			s.renderInventoryGui(entityManager, entityId)
		}
	}
}

func (s guiRenderSystem) renderCharacterStatsGui(entityManager *core.EntityManager, entityId string) {

	healthComponent, _ := entityManager.GetComponentOfClass(
		components.GetHealthComponentName(),
		entityId).(*components.Health)

	if healthComponent == nil { return }

	percentageHealth := float64(healthComponent.Health)/float64(healthComponent.MaxHealth)
	barWidth := HEALTH_BAR_MAX_WIDTH*percentageHealth

	ebitenutil.DrawRect(s.screen, HEALTH_BAR_POSITION_X, HEALTH_BAR_POSITION_Y, barWidth, HEALTH_BAR_HEIGHT, enum.RED)

	staminaComponent, _ := entityManager.GetComponentOfClass(
		components.GetStaminaComponentName(),
		entityId).(*components.Stamina)

	if staminaComponent == nil { return }

	percentageStamina := float64(staminaComponent.Stamina)/float64(staminaComponent.MaxStamina)
	barWidth = HEALTH_BAR_MAX_WIDTH*percentageStamina

	ebitenutil.DrawRect(s.screen, HEALTH_BAR_POSITION_X, HEALTH_BAR_POSITION_Y+HEALTH_BAR_HEIGHT, barWidth, HEALTH_BAR_HEIGHT, enum.GREEN)
}

func (s guiRenderSystem) renderInventoryGui(entityManager *core.EntityManager, entityId string) {

	inventoryComponent, _ := entityManager.GetComponentOfClass(
		components.GetInventoryComponentName(),
		entityId).(*components.Inventory)

	if inventoryComponent == nil { return }

	inputComponent, _ := entityManager.GetComponentOfClass(
		components.GetInputComponentName(),
		entityId).(*components.Input)

	if inputComponent == nil || inputComponent.InventoryToggle == false { return }

	inventoryFrame, _ := ebiten.NewImage(INVENTORY_WINDOW_WIDTH, INVENTORY_WINDOW_HEIGHT, ebiten.FilterNearest)
	inventoryFrame.Fill(enum.GREY)

	for idx, itemEntityId := range inventoryComponent.ItemEntityIds {
		if len(entityId) == 0 { continue }
		s.renderInventoryItems(entityManager, itemEntityId, inventoryFrame, idx)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(INVENTORY_WINDOW_POSITION_X), float64(INVENTORY_WINDOW_POSITION_Y))

	s.screen.DrawImage(inventoryFrame, op)
}

func (s guiRenderSystem) renderInventoryItems(
	entityManager *core.EntityManager,
	entityId string,
	inventoryFrame *ebiten.Image,
	inventoryIdx int,
) {

	spriteComponent, _ := entityManager.GetComponentOfClass(
		components.GetSpriteComponentName(),
		entityId).(*components.Sprite)

	if spriteComponent == nil { return }

	y := inventoryIdx/INVENTORY_MAX_CELLS_IN_ROW
	x := INVENTORY_MAX_CELLS_IN_ROW - (((y+1)*INVENTORY_MAX_CELLS_IN_ROW) - inventoryIdx)

	xOffset := x*INVENTORY_CELL_WIDTH
	yOffset := y*INVENTORY_CELL_HEIGHT

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.5,0.5)//this is a bit sketch

	op.GeoM.Translate(float64(xOffset), float64(yOffset))

	inventoryFrame.DrawImage(spriteComponent.ThumbnailImage, op)
}

func CreateGuiRenderSystem(screen *ebiten.Image) System {
	return guiRenderSystem{
		screen,
	}
}