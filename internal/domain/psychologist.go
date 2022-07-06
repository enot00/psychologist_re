package domain

type Psychologist struct {
	User           User
	Description    string
	Specialization []Specialization
}
