package world

import (
	"BasicECS/systems"
	"github.com/hajimehoshi/ebiten"
	"time"
	"fmt"
	"BasicECS/factories"
	"BasicECS/core"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/go-mplusbitmap"
	"BasicECS/enum"
)

var LAST = time.Now()
var SCREEN_WIDTH = 640
var SCREEN_HEIGHT = 480

type World struct {
	entityManager *core.EntityManager
	entityFactory *factories.EntityFactory
	systemManager *systems.SystemManager
	initialized bool
}

func (w *World) initializeWorld(screen *ebiten.Image) {
	playerEntityId := w.entityFactory.CreatePlayer(
		w.entityManager,
		550,
		240,
		2,
		100,
		400)

	w.entityFactory.CreateZombie(w.entityManager)

	w.entityFactory.CreateWeapon(
		w.entityManager,
		250,
		250,
		3,
		5,
		8)

	w.entityFactory.CreateMagazine(
		w.entityManager,
		350,
		350,
		30)

	w.entityFactory.CreateMagazine(
		w.entityManager,
		250,
		350,
		30)

	w.entityFactory.CreateMagazine(
		w.entityManager,
		150,
		350,
		30)

	w.entityFactory.CreateMagazine(
		w.entityManager,
		50,
		350,
		30)

	w.entityFactory.CreateGui(
		w.entityManager,
		playerEntityId,
	)

	w.systemManager.AddSystem(systems.CreateRemovalSystem())
	w.systemManager.AddSystem(systems.CreateExpirationSystem())
	w.systemManager.AddSystem(systems.CreateInputSystem())
	w.systemManager.AddSystem(systems.CreateMovementSystem())
	w.systemManager.AddSystem(systems.CreateCollisionSystem())
	w.systemManager.AddSystem(systems.CreateRenderSystem(screen))
	w.systemManager.AddSystem(systems.CreateInteractionSystem())
	w.systemManager.AddSystem(systems.CreateGuiRenderSystem(screen))
}

func (w *World) update(screen *ebiten.Image) error {

	if !w.initialized {
		w.initializeWorld(screen)
		w.initialized = true
	}

	now := time.Now()
	delta := float64(now.Sub(LAST))/float64(time.Second)
	LAST = now

	w.systemManager.ProcessSystems(delta, w.entityManager)

	fpsString := fmt.Sprintf("Current FPS: %f, delta: %f", ebiten.CurrentFPS(), delta)

	text.Draw(screen, fpsString, mplusbitmap.Gothic12r, 30, 10, enum.GREY)


	return nil
}

func (w *World) Run() {
	ebiten.Run(w.update, 640, 480, 2, "Hello world!")
}

func CreateWorld(
	entityManager *core.EntityManager,
	entityFactory *factories.EntityFactory,
	systemManager *systems.SystemManager) World {

	return World{
		entityManager,
		entityFactory,
		systemManager,
		false,
	}
}
