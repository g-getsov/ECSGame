package closures

import "BasicECS/core"

type InteractionFunc func(
	entityManager *core.EntityManager,
	entityId string,
	interactingEntityId string) bool

type UsageFunc func(entityManager core.EntityManager) bool