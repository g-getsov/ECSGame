package systems

import (
	"BasicECS/core"
)

type SystemManager struct {
	systems []System
}

func (sm *SystemManager) AddSystem(system System) {
	if system == nil { return }
	sm.systems = append(sm.systems, system)
	sm.sortSystems()
}

func (sm SystemManager) ProcessSystems(dt float64, entityManager *core.EntityManager) {
	for _, system := range sm.systems {
		system.Update(dt, entityManager)
	}
}

func (sm *SystemManager) sortSystems() {
	//TODO: Implement
}

func CreateSystemManager() SystemManager {
	return SystemManager {
		systems: make([]System, 0),
	}
}
