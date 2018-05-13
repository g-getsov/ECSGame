package systems

import (
	"BasicECS/core"
)

type System interface {
	Update(dt float64, entityManager *core.EntityManager)
}