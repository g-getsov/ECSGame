package world

import (
	"BasicECS/entities"
	"BasicECS/systems"
)

type World struct {
	entityManager *entities.EntityManager
	systemManager *systems.SystemManager
}

func (w World) Run() {
	for {
		w.systemManager.ProcessSystems(3.14, w.entityManager)
	}
}

func CreateWorld(entityManager *entities.EntityManager, systemManager *systems.SystemManager) World {
	return World{
		entityManager,
		systemManager,
	}
}
