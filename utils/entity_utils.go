package utils

import "BasicECS/core"

func RemoveEntity(entitySlice []core.Entity, entity core.Entity) []core.Entity {
	return RemoveEntityById(entitySlice, entity.Id)
}

func RemoveEntityById(entitySlice []core.Entity, entityId string) []core.Entity {
	for idx, value := range entitySlice {
		if value.Id == entityId {
			entitySlice = append(entitySlice[:idx], entitySlice[idx+1:]...)
			return entitySlice
		}
	}
	return entitySlice
}

func ContainsEntity(entitySlice []core.Entity, entity core.Entity) bool {
	return ContainsEntityById(entitySlice, entity.Id)
}

func ContainsEntityById(entitySlice []core.Entity, entityId string) bool {
	for _, value := range entitySlice {
		if value.Id == entityId {
			return true
		}
	}
	return false
}