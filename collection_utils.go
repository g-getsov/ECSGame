package main

import "BasicECS/entities"

func contains(intSlice []int, searchInt int) bool {
	for _, value := range intSlice {
		if value == searchInt {
			return true
		}
	}
	return false
}

func remove(intSlice []int, removeInt int) []int{
	intSlice = append(intSlice[:removeInt], intSlice[removeInt+1:]...)
	return intSlice
}

func containsEntity(entitySlice []entities.Entity, entity entities.Entity) bool {
	return containsEntityById(entitySlice, entity.Id)
}

func containsEntityById(entitySlice []entities.Entity, entityId int) bool {
	for _, value := range entitySlice {
		if value.Id == entityId {
			return true
		}
	}
	return false
}

func removeEntity(entitySlice []entities.Entity, entity entities.Entity) []entities.Entity {
	return removeEntityById(entitySlice, entity.Id)
}

func removeEntityById(entitySlice []entities.Entity, entityId int) []entities.Entity {
	for idx, value := range entitySlice {
		if value.Id == entityId {
			entitySlice = append(entitySlice[:idx], entitySlice[idx+1:]...)
			return entitySlice
		}
	}
	return entitySlice
}
