package components

const SprintingComponentName = "Sprinting"

type Sprinting struct {
	Id string
}

func (s Sprinting) GetComponentId() string { return s.Id }

func (s Sprinting) GetComponentName() string { return GetSprintingComponentName() }

func (s Sprinting) IsUniquePerEntity() bool { return true }

func GetSprintingComponentName() string {return SprintingComponentName }