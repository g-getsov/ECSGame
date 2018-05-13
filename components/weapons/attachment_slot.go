package weapons

import "BasicECS/enum"

const attachmentSlotComponentName = "AttachmentSlot"

type AttachmentSlot struct {
	EntityId string
	Type enum.AttachmentType
}

func (e AttachmentSlot) GetComponentName() string { return GetAttachmentSlotComponentName() }

func GetAttachmentSlotComponentName() string {return attachmentSlotComponentName }