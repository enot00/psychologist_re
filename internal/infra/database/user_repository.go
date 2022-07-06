package database

import (
	"github.com/test_server/internal/domain"
	"github.com/upper/db/v4"
	"log"
	"time"
)

type user struct {
	ID         int64           `db:"id,omitempty"`
	Phone      string          `db:"phone,omitempty"`
	Email      string          `db:"email,omitempty"`
	Password   string          `db:"password,omitempty"`
	FirstName  string          `db:"first_name,omitempty"`
	SecondName *string         `db:"second_name,omitempty"`
	LastName   string          `db:"last_name,omitempty"`
	Avatar     string          `db:"avatar,omitempty"`
	Role       domain.UserRole `db:"role,omitempty"`
	CreatedAt  time.Time       `db:"created_at,omitempty"`
	DeletedAt  time.Time       `db:"deleted_at,omitempty"`
}

type userShort struct {
	ID         int64   `db:"id,omitempty"`
	FirstName  string  `db:"first_name"`
	SecondName *string `db:"second_name,omitempty"`
	LastName   string  `db:"last_name"`
	Avatar     string  `db:"avatar,omitempty"`
}

type UserRepository interface {
	PaginateAll(page, pageSize uint) (*[]domain.UserShort, error)
	GetOneByID(id int64) (*domain.User, error)
	GetOneByEmail(email string) (*domain.User, error)
	Save(user *domain.User) (*domain.User, error)
	Update(user *domain.User) error
	Delete(userID int64) error
}

type userRepository struct {
	collDB   db.Session
	userColl db.Collection
}

func NewUserRepository(dbSession *db.Session) UserRepository {
	return &userRepository{
		collDB:   *dbSession,
		userColl: (*dbSession).Collection("users"),
	}
}

func (u *userRepository) PaginateAll(page, pageSize uint) (*[]domain.UserShort, error) {

	var users []userShort

	err := u.userColl.Find().Paginate(pageSize).Page(page).All(&users)

	if err != nil {
		log.Printf("userRepository.PaginateAll(): %s", err.Error())
		return nil, err
	}

	return mapUserDbModelToDomainCollection(&users), nil
}

func (u *userRepository) GetOneByID(id int64) (*domain.User, error) {

	var userOne user
	err := u.userColl.Find(db.Cond{"id": id}).One(&userOne)
	if err != nil {
		log.Printf("userRepository.GetOneByID(): %s", err.Error())
		return nil, err
	}

	return mapUserDbModelToDomain(&userOne), nil
}
func (u *userRepository) GetOneByEmail(email string) (*domain.User, error) {

	var ps user
	err := u.userColl.Find(db.Cond{"email": email}).One(&ps)
	if err != nil {
		log.Printf("userRepository.GetOneByEmail(): %s", err.Error())
		return nil, err
	}

	return mapUserDbModelToDomain(&ps), nil
}

func (u *userRepository) Save(user *domain.User) (*domain.User, error) {
	userWithID := mapUserDomainToDbModel(user)
	err := u.userColl.InsertReturning(userWithID)
	if err != nil {
		log.Printf("userRepository.Save(): %s", err.Error())
		return nil, err
	}
	userDomain := mapUserDbModelToDomain(userWithID)
	return userDomain, nil
}

func (u *userRepository) Update(user *domain.User) error {
	userDomain := mapUserDomainToDbModel(user)
	err := u.userColl.Find(db.Cond{"id": user.ID}).Update(userDomain)
	if err != nil {
		log.Printf("userRepository.Update(): %s", err.Error())
		return err
	}
	return nil
}
func (u *userRepository) Delete(userID int64) error {

	now := time.Now().Local()

	q := u.collDB.SQL().Update("users").Set("deleted_at = ?", now).Where("id = ?", userID)
	_, err := q.Exec()
	if err != nil {
		log.Printf("userRepository.Delete(): %s", err.Error())
		return err
	}

	return nil
}

func mapUserDbModelToDomain(us *user) *domain.User {
	return &domain.User{
		ID:         us.ID,
		Phone:      us.Phone,
		Email:      us.Email,
		Password:   us.Password,
		FirstName:  us.FirstName,
		SecondName: us.SecondName,
		LastName:   us.LastName,
		Avatar:     us.Avatar,
		Role:       us.Role,
		CreatedAt:  us.CreatedAt,
		DeletedAt:  us.DeletedAt,
	}
}

func mapUserDbModelToDomainCollection(u *[]userShort) *[]domain.UserShort {
	var result []domain.UserShort

	for _, v := range *u {
		newPS := mapUserShortDomainToDbModel(&v)
		result = append(result, *newPS)
	}

	return &result
}

func mapUserDomainToDbModel(us *domain.User) *user {

	return &user{
		ID:         us.ID,
		Phone:      us.Phone,
		Email:      us.Email,
		Password:   us.Password,
		FirstName:  us.FirstName,
		SecondName: us.SecondName,
		LastName:   us.LastName,
		Avatar:     us.Avatar,
		Role:       us.Role,
		CreatedAt:  us.CreatedAt,
		DeletedAt:  us.DeletedAt,
	}
}
func mapUserShortDomainToDbModel(us *userShort) *domain.UserShort {

	return &domain.UserShort{
		ID:         us.ID,
		FirstName:  us.FirstName,
		SecondName: us.SecondName,
		LastName:   us.LastName,
		Avatar:     us.Avatar,
	}
}
