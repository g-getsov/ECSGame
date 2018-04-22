package entities

import (
	"math"
	"fmt"
	cmpt "BasicECS/components"
)

type EntityManager struct {
	lowestUnassignedEntityId int
	entities []Entity
	componentsByClass map[string]map[int]cmpt.Component
}

func (e *EntityManager) generateNewEntityId() int {
	if e.lowestUnassignedEntityId < math.MaxInt32 {
		e.lowestUnassignedEntityId++
		return e.lowestUnassignedEntityId
	}

	for i := 0; i < math.MaxInt32; i++ {
		if ContainsEntityById(e.entities, i) {
			return i
		}
	}

	fmt.Println("No available entity ids")
	return 0
}

func (e *EntityManager) CreateEntity() Entity {
	id := e.generateNewEntityId()
	entity := CreateNewEntity(id)
	e.entities = append(e.entities, entity)
	return entity
}

func (e *EntityManager) RemoveEntity(entityId int) {
	for _, component := range e.componentsByClass {
		if component[entityId] != nil {
			delete(component, entityId)
		}
	}
	e.entities = RemoveEntityById(e.entities, entityId)
}

func (e *EntityManager) AddComponentToEntity(entityId int, component cmpt.Component) {
	e.AddComponentsToEntity(entityId, []cmpt.Component{component})
}

func (e *EntityManager) AddComponentsToEntity(entityId int, components []cmpt.Component) {

	for _, component := range components {

		componentName := component.GetComponentName()

		componentsForClass := e.componentsByClass[componentName]

		if componentsForClass == nil {
			componentsForClass := make(map[int]cmpt.Component)
			componentsForClass[entityId] = component
			e.componentsByClass[componentName] = componentsForClass
			continue
		}

		componentsForClass[entityId] = component
	}
}

func (e *EntityManager) GetComponentOfClass(componentName string, entityId int) cmpt.Component {
	componentsForClass := e.componentsByClass[componentName]
	if componentsForClass == nil { return nil }
	return componentsForClass[entityId]
}
func (e *EntityManager) GetAllEntitiesPossessingComponentsOfClass(componentName string) []int {
	componentsForClass := e.componentsByClass[componentName]
	if componentsForClass == nil { return make([]int, 0)}

	entitiesPossessingComponent := make([]int, len(componentsForClass))
	i := 0
	for key := range componentsForClass {
		entitiesPossessingComponent[i] = key
		i++
	}
	return entitiesPossessingComponent
}

func CreateEntityManager() EntityManager {
	return EntityManager {
		lowestUnassignedEntityId: 0,
		entities: make([]Entity, 0),
		componentsByClass: make(map[string]map[int]cmpt.Component),
	}
}