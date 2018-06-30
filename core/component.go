package core

type Component interface {
	GetComponentId() string
	GetComponentName() string
	IsUniquePerEntity() bool
}
