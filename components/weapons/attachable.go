package weapons

import "BasicECS/enum"

const attachableComponentName = "Attachable"

type Attachable struct {
	AttachmentType enum.AttachmentType
}

func (e Attachable) GetComponentName() string { return GetAttachableComponentName() }

func GetAttachableComponentName() string {return attachableComponentName }