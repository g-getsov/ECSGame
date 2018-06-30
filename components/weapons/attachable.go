package weapons

import "BasicECS/enum"

const attachableComponentName = "Attachable"

type Attachable struct {
	Id string
	AttachmentType enum.AttachmentType
}

func (a Attachable) GetComponentId() string { return a.Id }

func (a Attachable) GetComponentName() string { return GetAttachableComponentName() }

func (a Attachable) IsUniquePerEntity() bool { return true }

func GetAttachableComponentName() string {return attachableComponentName }