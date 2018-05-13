package main

import (
	"BasicECS/world"
	"BasicECS/systems"
	"BasicECS/resources"
	"BasicECS/factories"
	"BasicECS/core"
)

func prepareWorld() world.World {
	entityManager := core.CreateEntityManager(500)
	systemManager := systems.CreateSystemManager()
	imageManager := resources.CreateImageManager()
	entityFactory := factories.CreateEntityFactory(&imageManager)
	return world.CreateWorld(&entityManager, &entityFactory, &systemManager)
}

func main() {
	newWorld := prepareWorld()
	newWorld.Run()
}