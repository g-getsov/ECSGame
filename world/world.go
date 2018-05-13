package world

import (
	"BasicECS/entities"
	"BasicECS/systems"
	"github.com/hajimehoshi/ebiten"
	"time"
	"fmt"
)

var last = time.Now()

type World struct {
	entityManager *entities.EntityManager
	entityFactory *entities.EntityFactory
	systemManager *systems.SystemManager
	initialized bool
}

func (w *World) initializeWorld(screen *ebiten.Image) {
	w.entityFactory.CreatePlayer(
		w.entityManager,
		550,
		240,
		2,
		100)

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

	w.systemManager.AddSystem(systems.CreateInputSystem())
	w.systemManager.AddSystem(systems.CreateMovementSystem())
	w.systemManager.AddSystem(systems.CreateRenderSystem(screen))
}

func (w *World) update(screen *ebiten.Image) error {

	if !w.initialized {
		w.initializeWorld(screen)
		w.initialized = true
	}

	now := time.Now()
	delta := float64(now.Sub(last))/float64(time.Second)
	last = now

	w.systemManager.ProcessSystems(delta, w.entityManager)
	fmt.Println("Current FPS:", ebiten.CurrentFPS())

	return nil
}

func (w *World) Run() {
	ebiten.Run(w.update, 640, 480, 2, "Hello world!")
}

func CreateWorld(
	entityManager *entities.EntityManager,
	entityFactory *entities.EntityFactory,
	systemManager *systems.SystemManager) World {

	return World{
		entityManager,
		entityFactory,
		systemManager,
		false,
	}
}
