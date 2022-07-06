package domain

type Chat struct {
	Id           uint64
	CreatorId    uint64
	Participants []uint64
}
