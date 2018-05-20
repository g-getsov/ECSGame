package utils

import (
	"BasicECS/core"
	"BasicECS/components"
)

func IsEntityHidden(entityManager *core.EntityManager, entityId string) bool {
	return entityManager.GetComponentOfClass(components.GetHiddenComponentName(), entityId) != nil
}