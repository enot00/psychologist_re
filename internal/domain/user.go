package domain

import "time"

type UserRole uint8

const AvatarPath = "/psychology-go-back/file_storage/user_avatar/"

const (
	ClientRole UserRole = iota
	PsychologistRole
)

type User struct {
	ID         int64
	Phone      string
	Email      string
	Password   string
	FirstName  string
	SecondName *string
	LastName   string
	Avatar     string
	Role       UserRole
	CreatedAt  time.Time
	DeletedAt  time.Time
}
type UserShort struct {
	ID         int64
	FirstName  string
	SecondName *string
	LastName   string
	Avatar     string
}

func (u *User) IsDeleted() bool {
	return !u.DeletedAt.IsZero()
}
func (us *UserShort) GetAvatarShort() string {
	if us.Avatar != "" {
		return AvatarPath + us.Avatar
	} else {
		return us.Avatar
	}
}
func (u *User) GetAvatar() string {
	if u.Avatar != "" {
		return AvatarPath + u.Avatar
	} else {
		return u.Avatar
	}
}
