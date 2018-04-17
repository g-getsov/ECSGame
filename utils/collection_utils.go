package utils

import "BasicECS/entities"

func Contains(intSlice []int, searchInt int) bool {
	for _, value := range intSlice {
		if value == searchInt {
			return true
		}
	}
	return false
}

func Remove(intSlice []int, removeInt int) []int{
	intSlice = append(intSlice[:removeInt], intSlice[removeInt+1:]...)
	return intSlice
}

func ContainsEntity(entitySlice []entities.Entity, entity entities.Entity) bool {
	return ContainsEntityById(entitySlice, entity.Id)
}

func ContainsEntityById(entitySlice []entities.Entity, entityId int) bool {
	for _, value := range entitySlice {
		if value.Id == entityId {
			return true
		}
	}
	return false
}

func RemoveEntity(entitySlice []entities.Entity, entity entities.Entity) []entities.Entity {
	return RemoveEntityById(entitySlice, entity.Id)
}

func RemoveEntityById(entitySlice []entities.Entity, entityId int) []entities.Entity {
	for idx, value := range entitySlice {
		if value.Id == entityId {
			entitySlice = append(entitySlice[:idx], entitySlice[idx+1:]...)
			return entitySlice
		}
	}
	return entitySlice
}
