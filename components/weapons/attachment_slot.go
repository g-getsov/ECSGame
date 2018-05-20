package weapons

import "BasicECS/enum"

const attachmentSlotComponentName = "AttachmentSlot"

type AttachmentSlot struct {
	EntityId string
	Type enum.AttachmentType
}

func (a AttachmentSlot) GetComponentName() string { return GetAttachmentSlotComponentName() }

func (a AttachmentSlot) IsUniquePerEntity() bool { return true }

func GetAttachmentSlotComponentName() string {return attachmentSlotComponentName }