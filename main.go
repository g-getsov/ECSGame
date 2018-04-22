package main

import (
	"BasicECS/world"
	"BasicECS/entities"
	"BasicECS/systems"
	"BasicECS/resources"
)

func prepareWorld() world.World {
	entityManager := entities.CreateEntityManager()
	systemManager := systems.CreateSystemManager()
	imageManager := resources.CreateImageManager()
	entityFactory := entities.CreateEntityFactory(&imageManager)
	return world.CreateWorld(&entityManager, &entityFactory, &systemManager)
}

func main() {
	newWorld := prepareWorld()
	newWorld.Run()
}