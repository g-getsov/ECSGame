package core

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
	"BasicECS/factories"
	"BasicECS/utils/uidutils"
)

func TestNewIdGeneration(t *testing.T) {

	id := uidutils.GenerateNewUniqueId()
	secondId := uidutils.GenerateNewUniqueId()

	assert.True(t, len(id) > 0)
	assert.True(t, len(secondId) > 0)
	assert.True(t, strings.Compare(id, secondId) != 0)
}

func TestFullIdGeneration(t *testing.T) {

	id := uidutils.GenerateNewUniqueId()
	secondId := uidutils.GenerateNewUniqueId()

	assert.True(t, len(id) > 0)
	assert.True(t, len(secondId) > 0)
}

func TestCreateEntities(t *testing.T) {

	entityManager := CreateEntityManager(4)

	entity := entityManager.CreateEntity()
	secondEntity := entityManager.CreateEntity()

	assert.True(t, len(entity.Id) > 0)
	_, ok := entityManager.entities[entity.Id]
	assert.True(t, ok)
	assert.True(t, len(secondEntity.Id) > 0)
	_, ok = entityManager.entities[secondEntity.Id]
	assert.True(t, ok)
}

func TestCreateEntitiesFullCapacity(t *testing.T) {

	entityManager := CreateEntityManager(1)

	entity := entityManager.CreateEntity()
	secondEntity := entityManager.CreateEntity()

	assert.True(t, len(entity.Id) > 0)
	_, ok := entityManager.entities[entity.Id]
	assert.True(t, ok)

	assert.True(t, strings.Compare(secondEntity.Id, "") == 0)
	_, ok = entityManager.entities[secondEntity.Id]
	assert.False(t, ok)

	entityManager.RemoveEntity(entity.Id)

	secondEntity = entityManager.CreateEntity()

	assert.True(t, len(secondEntity.Id) > 0)
	_, ok = entityManager.entities[secondEntity.Id]
	assert.True(t, ok)
}

func TestAddComponentToEntity(t *testing.T) {

	entityManager := CreateEntityManager(4)

	entity := entityManager.CreateEntity()
	secondEntity := entityManager.CreateEntity()

	speedComponent := factories.CreateSpeedComponent(32)
	healthComponent := factories.CreateHealthComponent(10)

	healthComponentName := healthComponent.GetComponentName()
	speedComponentName := speedComponent.GetComponentName()

	assert.Nil(t, entityManager.componentsByClass[speedComponentName])
	assert.Nil(t, entityManager.componentsByClass[healthComponentName])

	entityManager.AddComponentToEntity(entity.Id, speedComponent)
	entityManager.AddComponentToEntity(secondEntity.Id, healthComponent)

	assert.NotNil(t, entityManager.componentsByClass[speedComponentName])
	assert.NotNil(t, entityManager.componentsByClass[healthComponentName])

	assert.Nil(t, entityManager.componentsByClass[healthComponentName][entity.Id])
	assert.NotNil(t, entityManager.componentsByClass[speedComponentName][entity.Id])

	assert.NotNil(t, entityManager.componentsByClass[healthComponentName][secondEntity.Id])

	assert.Nil(t, entityManager.componentsByClass["invalidcomponent"])
}

func TestAddComponentsToEntity(t *testing.T)  {

	entityManager := CreateEntityManager(4)

	entity := entityManager.CreateEntity()
	secondEntity := entityManager.CreateEntity()

	speedComponent := factories.CreateSpeedComponent(32)
	healthComponent := factories.CreateHealthComponent(10)

	healthComponentName := healthComponent.GetComponentName()
	speedComponentName := speedComponent.GetComponentName()

	assert.Nil(t, entityManager.componentsByClass[speedComponentName])
	assert.Nil(t, entityManager.componentsByClass[healthComponentName])

	entityManager.AddComponentsToEntity(entity.Id, []Component{speedComponent, healthComponent})
	entityManager.AddComponentsToEntity(secondEntity.Id, []Component{healthComponent})

	assert.NotNil(t, entityManager.componentsByClass[speedComponentName])
	assert.NotNil(t, entityManager.componentsByClass[healthComponentName])

	assert.NotNil(t, entityManager.componentsByClass[healthComponentName][entity.Id])
	assert.NotNil(t, entityManager.componentsByClass[speedComponentName][entity.Id])

	assert.NotNil(t, entityManager.componentsByClass[healthComponentName][secondEntity.Id])

	assert.Nil(t, entityManager.componentsByClass["invalidcomponent"])
}

func TestGetComponentOfClass(t *testing.T) {

	entityManager := CreateEntityManager(4)

	entity := entityManager.CreateEntity()
	secondEntity := entityManager.CreateEntity()

	speedComponent := factories.CreateSpeedComponent(32)
	healthComponent := factories.CreateHealthComponent(10)

	healthComponentName := healthComponent.GetComponentName()
	speedComponentName := speedComponent.GetComponentName()

	assert.Nil(t, entityManager.componentsByClass[speedComponentName])
	assert.Nil(t, entityManager.componentsByClass[healthComponentName])

	entityManager.AddComponentToEntity(entity.Id, speedComponent)
	entityManager.AddComponentToEntity(secondEntity.Id, healthComponent)

	storedSpeedComponent := entityManager.GetComponentOfClass(speedComponentName, entity.Id)
	storedHealthComponent := entityManager.GetComponentOfClass(healthComponentName, secondEntity.Id)
	noComponent := entityManager.GetComponentOfClass(healthComponentName, entity.Id)

	assert.NotNil(t, storedSpeedComponent)
	assert.NotNil(t, storedHealthComponent)
	assert.Equal(t, healthComponent, storedHealthComponent)
	assert.Equal(t, speedComponent, storedSpeedComponent)
	assert.Nil(t, noComponent)
}

func TestRemoveEntity(t *testing.T) {

	entityManager := CreateEntityManager(4)

	entity := entityManager.CreateEntity()
	secondEntity := entityManager.CreateEntity()

	speedComponent := factories.CreateSpeedComponent(32)

	speedComponentName := speedComponent.GetComponentName()

	assert.Nil(t, entityManager.componentsByClass[speedComponentName])

	assert.Nil(t, entityManager.GetComponentOfClass(speedComponentName, entity.Id))
	assert.Nil(t, entityManager.GetComponentOfClass(speedComponentName, secondEntity.Id))

	entityManager.AddComponentToEntity(entity.Id, speedComponent)
	entityManager.AddComponentToEntity(secondEntity.Id, speedComponent)

	entityManager.RemoveEntity(entity.Id)

	returnedSpeedComponent := entityManager.GetComponentOfClass(speedComponentName, entity.Id)
	assert.Nil(t, returnedSpeedComponent)

	_, ok := entityManager.entities[entity.Id]
	assert.False(t, ok)

	secondReturnedSpeedComponent := entityManager.GetComponentOfClass(speedComponentName, secondEntity.Id)
	assert.NotNil(t, secondReturnedSpeedComponent)
}

func TestGetAllEntitiesPossessingComponentsOfClass(t *testing.T) {

	entityManager := CreateEntityManager(4)

	entity := entityManager.CreateEntity()
	secondEntity := entityManager.CreateEntity()

	speedComponent := factories.CreateSpeedComponent(32)
	healthComponent := factories.CreateHealthComponent(10)

	healthComponentName := healthComponent.GetComponentName()
	speedComponentName := speedComponent.GetComponentName()

	entityManager.AddComponentToEntity(entity.Id, speedComponent)
	entityManager.AddComponentToEntity(entity.Id, healthComponent)
	entityManager.AddComponentToEntity(secondEntity.Id, healthComponent)

	healthEntities := entityManager.GetAllEntitiesPossessingComponentsOfClass(healthComponentName)
	speedEntities := entityManager.GetAllEntitiesPossessingComponentsOfClass(speedComponentName)
	positionEntities := entityManager.GetAllEntitiesPossessingComponentsOfClass("Position")

	assert.Equal(t, 2, len(healthEntities))
	assert.Equal(t, 1, len(speedEntities))
	assert.Equal(t, 0, len(positionEntities))
}

