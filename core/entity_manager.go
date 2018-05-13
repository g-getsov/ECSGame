package core

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type EntityManager struct {
	maxNumEntities int
	entities map[string]Entity
	componentsByClass map[string]map[string]Component
}

func (e *EntityManager) generateNewEntityId() string {

	id, err := uuid.NewV4()

	if err != nil {
		fmt.Println("Something went wrong while generating a new entity id")
		return ""
	}

	return id.String()
}

func (e *EntityManager) CreateEntity() Entity {
	id := e.generateNewEntityId()

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

func (e *EntityManager) AddComponentToEntity(entityId string, component Component) {
	e.AddComponentsToEntity(entityId, []Component{component})
}

func (e *EntityManager) AddComponentsToEntity(entityId string, components []Component) {

	for _, component := range components {

		componentName := component.GetComponentName()

		componentsForClass := e.componentsByClass[componentName]

		if componentsForClass == nil {
			componentsForClass := make(map[string]Component)
			componentsForClass[entityId] = component
			e.componentsByClass[componentName] = componentsForClass
			continue
		}

		componentsForClass[entityId] = component
	}
}

func (e *EntityManager) GetComponentOfClass(componentName string, entityId string) Component {
	componentsForClass := e.componentsByClass[componentName]
	if componentsForClass == nil { return nil }
	return componentsForClass[entityId]
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
		componentsByClass: make(map[string]map[string]Component),
	}
}