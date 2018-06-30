package components

const tooltipComponentName = "Tooltip"

type Tooltip struct {
	Id string
	Text string
}

func (t Tooltip) GetComponentId() string { return t.Id }

func (t Tooltip) GetComponentName() string { return GetTooltipComponentName() }

func (t Tooltip) IsUniquePerEntity() bool { return true }

func GetTooltipComponentName() string {return tooltipComponentName }