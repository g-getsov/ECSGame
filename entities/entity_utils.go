package entities

func RemoveEntity(entitySlice []Entity, entity Entity) []Entity {
	return RemoveEntityById(entitySlice, entity.Id)
}

func RemoveEntityById(entitySlice []Entity, entityId string) []Entity {
	for idx, value := range entitySlice {
		if value.Id == entityId {
			entitySlice = append(entitySlice[:idx], entitySlice[idx+1:]...)
			return entitySlice
		}
	}
	return entitySlice
}

func ContainsEntity(entitySlice []Entity, entity Entity) bool {
	return ContainsEntityById(entitySlice, entity.Id)
}

func ContainsEntityById(entitySlice []Entity, entityId string) bool {
	for _, value := range entitySlice {
		if value.Id == entityId {
			return true
		}
	}
	return false
}