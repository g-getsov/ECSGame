package components

const tooltipComponentName = "Tooltip"

type Tooltip struct {
	Text string
}

func (t Tooltip) GetComponentName() string { return GetTooltipComponentName() }

func (t Tooltip) IsUniquePerEntity() bool { return true }

func GetTooltipComponentName() string {return tooltipComponentName }