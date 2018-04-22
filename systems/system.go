package systems

import "BasicECS/entities"

type System interface {
	Update(dt float64, entityManager *entities.EntityManager)
}