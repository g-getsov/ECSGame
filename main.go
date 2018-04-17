package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"BasicECS/world"
	"BasicECS/entities"
	"BasicECS/systems"
)

func update(screen *ebiten.Image) error {
	ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}

func prepareWorld() world.World {
	entityManager := entities.CreateEntityManager()
	systemManager := systems.CreateSystemManager()
	return world.CreateWorld(&entityManager, &systemManager)
}

func main() {
/*	newWorld := prepareWorld()
	newWorld.Run()*/
	ebiten.Run(update, 320, 240, 2, "Hello world!")
}