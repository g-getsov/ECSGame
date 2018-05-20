package components

const velocityComponentName = "Velocity"

type Velocity struct {
	X int
	Y int
}

func (v Velocity) GetComponentName() string { return GetVelocityComponentName() }

func (v Velocity) IsUniquePerEntity() bool { return true }

func GetVelocityComponentName() string { return velocityComponentName }