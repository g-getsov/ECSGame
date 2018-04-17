package systems

import "BasicECS/entities"

type System interface {
	Update(dt float32, entityManager *entities.EntityManager)
}