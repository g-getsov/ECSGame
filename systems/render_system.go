package systems

import "BasicECS/entities"

type renderSystem struct {}

func (s renderSystem) Update(dt float32, entityManager *entities.EntityManager) {

}

func CreateRenderSystem() System { return renderSystem{} }
