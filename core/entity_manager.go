package core

import (
	"fmt"
	"BasicECS/utils/uidutils"
)

type EntityManager struct {
	maxNumEntities int
	entities map[string]Entity
	componentsByClass map[string]map[string]map[string]Component //componentName -> entityId -> componentId
}

func (e *EntityManager) CreateEntity() Entity {
	id := uidutils.GenerateNewUniqueId()

	var entity Entity

	if len(e.entities) >= e.maxNumEntities {
		fmt.Println("Too many entities")
		return entity
	}

	entity = CreateNewEntity(id)
	e.entities[id] = entity
	return entity
}

func (e *EntityManager) RemoveEntity(entityId string) {
	for _, component := range e.componentsByClass {
		if component[entityId] != nil {
			delete(component, entityId)
		}
	}
	delete(e.entities, entityId)
}

func (e *EntityManager) RemoveComponentForEntity(entityId string, componentName string) {
	componentsForClass := e.componentsByClass[componentName]
	if componentsForClass != nil {
		delete(componentsForClass, entityId)
	}
}

func (e *EntityManager) AddComponentToEntity(entityId string, component Component) {
	e.AddComponentsToEntity(entityId, []Component{component})
}

func (e *EntityManager) AddComponentsToEntity(entityId string, components []Component) {

	for _, component := range components {

		componentName := component.GetComponentName()

		componentsForClass := e.componentsByClass[componentName]

		if componentsForClass == nil {
			componentsForClass := make(map[string]map[string]Component)
			e.componentsByClass[componentName] = componentsForClass
		}

		// if component is unique disassociate all existing components
		if component.IsUniquePerEntity() {
			e.componentsByClass[componentName][entityId] = make(map[string]Component)
		}

		if e.componentsByClass[componentName][entityId] == nil {
			e.componentsByClass[componentName][entityId] = make(map[string]Component)
		}

		e.componentsByClass[componentName][entityId][component.GetComponentId()] = component
	}
}

func (e *EntityManager) GetComponentOfClass(componentName string, entityId string) Component {

	componentsForClass := e.componentsByClass[componentName]

	if componentsForClass == nil { return nil }

	componentsForEntity := componentsForClass[entityId]

	if componentsForEntity == nil { return nil }

	for _, component := range componentsForEntity {
		if !component.IsUniquePerEntity() {
			panic(fmt.Sprintf(
				"Trying to retrieve non-unique component of class %s as unique !",
				component.GetComponentName()))
		}
		return component
	}

	return nil
}

func (e *EntityManager) GetComponent(componentName string, entityId string, componentId string) Component {

	componentsForClass := e.componentsByClass[componentName]

	if componentsForClass == nil { return nil }

	componentsForEntity := componentsForClass[entityId]

	if componentsForEntity == nil { return nil }

	return componentsForEntity[componentId]
}

func (e *EntityManager) GetComponentsOfClass(componentName string, entityId string) []Component {

	componentsForClass := e.componentsByClass[componentName]

	if componentsForClass == nil { return nil }

	componentsForEntity := componentsForClass[entityId]

	if componentsForEntity == nil { return nil }

	components := make([]Component, len(componentsForEntity))

	for _, component := range componentsForEntity {
		components = append(components, component)
	}

	return components
}

func (e *EntityManager) GetAllEntitiesPossessingComponentsOfClass(componentName string) []string {
	componentsForClass := e.componentsByClass[componentName]
	if componentsForClass == nil { return make([]string, 0)}

	entitiesPossessingComponent := make([]string, len(componentsForClass))
	i := 0
	for key := range componentsForClass {
		entitiesPossessingComponent[i] = key
		i++
	}
	return entitiesPossessingComponent
}

func CreateEntityManager(maxNumEntities int) EntityManager {
	return EntityManager {
		maxNumEntities: maxNumEntities,
		entities: make(map[string]Entity),
		componentsByClass: make(map[string]map[string]map[string]Component),
	}
}