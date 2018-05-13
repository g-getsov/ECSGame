package closures

import "BasicECS/entities"

type InteractionFunc func(
	entityManager *entities.EntityManager,
	entityId int,
	interactingEntityId int) bool