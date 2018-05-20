package core

type Component interface {
	GetComponentName() string
	IsUniquePerEntity() bool
}
