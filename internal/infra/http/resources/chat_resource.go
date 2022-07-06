package resources

type ChatDto struct {
	Id           uint64   `json:"id,omitempty"`
	CreatorId    uint64   `json:"creator_id" validate:"required"`
	Participants []uint64 `json:"participants" validate:"required,min=2"`
}
