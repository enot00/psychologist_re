package database

import (
	"github.com/test_server/internal/domain"
	"github.com/upper/db/v4"
)

type psychologist struct {
	User           user             `db:",inline"`
	Description    string           `db:"description"`
	Specialization []specialization `db:",inline"`
}
type specialization struct {
	ID    int64  `db:"id,omitempty"`
	Name  string `db:"name"`
	PsyID int64  `db:"psy_id"`
}

type psychologistsSpecializations struct {
	PsychologistID   int64 `db:"psychologist_id"`
	SpecializationID int64 `db:"specialization_id"`
}

type PsychologistRepository interface {
	PaginateAll(page, pageSize uint) (*[]domain.Psychologist, error)
	GetOne(id uint) (*domain.Psychologist, error)
	Update(psychologist *domain.Psychologist) error
}

type psychologistRepository struct {
	coll     db.Session
	psyColl  db.Collection
	specColl db.Collection
}

func NewClientPsychologistRepository(dbSession *db.Session) PsychologistRepository {
	return &psychologistRepository{
		coll:     *dbSession,
		psyColl:  (*dbSession).Collection("users"),
		specColl: (*dbSession).Collection("users_specializations"),
	}
}

func (r *psychologistRepository) PaginateAll(page, pageSize uint) (*[]domain.Psychologist, error) {
	var userShort []userShort
	err := r.psyColl.Find(db.Cond{"role": domain.PsychologistRole}).Paginate(pageSize).Page(page).All(&userShort)
	if err != nil {
		return nil, err
	}
	var psyID []int64
	for _, v := range userShort {
		psyID = append(psyID, v.ID)
	}
	var specializations []specialization
	querySpec := r.coll.SQL().SelectFrom("specializations s").
		Columns("s.*", "us.psychologist_id AS psy_id").
		Join("users_specializations AS us").
		On("s.id = us.specialization_id").Where(db.Cond{"us.psychologist_id IN": psyID})
	err = querySpec.All(&specializations)
	return mapPsychologistDbModelToDomainCollection(&userShort, &specializations), nil
}

func (r *psychologistRepository) GetOne(id uint) (*domain.Psychologist, error) {
	var ps psychologist
	err := r.psyColl.Find(db.Cond{"id": id}).One(&ps)
	if err != nil {
		return nil, err
	}
	var specializations []specialization
	var psyID []uint
	psyID = append(psyID, id)

	if ps.User.Role == domain.PsychologistRole {
		querySpec := r.coll.SQL().SelectFrom("specializations s").
			Columns("s.*", "us.psychologist_id AS psy_id").
			Join("users_specializations AS us").
			On("s.id = us.specialization_id").Where(db.Cond{"us.psychologist_id": ps.User.ID})
		err = querySpec.All(&specializations)
	}
	return mapPsychologistDbModelToDomain(&ps, &specializations), nil
}

func (r *psychologistRepository) Update(psychologist *domain.Psychologist) error {
	psy := mapPsychologistDomainToDbModel(psychologist)
	err := r.psyColl.Find(db.Cond{"id": psy.User.ID}).Update(psy)
	if err != nil {
		return err
	}
	spec := mapSpecializationDomainToDbModel(&psychologist.Specialization)
	err = r.specColl.Find(db.Cond{"psychologist_id": psy.User.ID}).Delete()
	if err != nil {
		return err
	}
	for _, v := range *spec {
		ps := psychologistsSpecializations{
			PsychologistID:   psy.User.ID,
			SpecializationID: v.ID,
		}
		err = r.specColl.InsertReturning(&ps)
		if err != nil {
			return err
		}
	}
	return nil
}

func mapPsychologistShortDbModelToDomain(ps *userShort, spec *[]specialization) *domain.Psychologist {
	sp := mapSpecializationToPsychologistDbModel(ps.ID, spec)

	return &domain.Psychologist{
		User: domain.User{
			ID:         ps.ID,
			FirstName:  ps.FirstName,
			SecondName: ps.SecondName,
			LastName:   ps.LastName,
			Avatar:     ps.Avatar,
		},
		Specialization: *sp,
	}
}
func mapPsychologistDbModelToDomain(ps *psychologist, spec *[]specialization) *domain.Psychologist {
	sp := mapSpecializationToPsychologistDbModel(ps.User.ID, spec)

	return &domain.Psychologist{
		User: domain.User{
			ID:         ps.User.ID,
			Phone:      ps.User.Phone,
			Email:      ps.User.Email,
			Password:   ps.User.Password,
			FirstName:  ps.User.FirstName,
			SecondName: ps.User.SecondName,
			LastName:   ps.User.LastName,
			Avatar:     ps.User.Avatar,
			Role:       ps.User.Role,
			CreatedAt:  ps.User.CreatedAt,
			DeletedAt:  ps.User.DeletedAt,
		},
		Description:    ps.Description,
		Specialization: *sp,
	}
}

func mapPsychologistDbModelToDomainCollection(ps *[]userShort, sp *[]specialization) *[]domain.Psychologist {
	var result []domain.Psychologist

	for _, v := range *ps {
		newPS := mapPsychologistShortDbModelToDomain(&v, sp)
		result = append(result, *newPS)
	}

	return &result
}
func mapSpecializationToPsychologistDbModel(id int64, sp *[]specialization) *[]domain.Specialization {
	var result []domain.Specialization
	for _, v := range *sp {
		if id == v.PsyID {
			result = append(result, domain.Specialization{ID: v.ID, Name: v.Name})
		}
	}
	return &result
}

func mapPsychologistDomainToDbModel(ps *domain.Psychologist) *psychologist {
	spec := mapSpecializationDomainToDbModel(&ps.Specialization)

	return &psychologist{
		User: user{
			ID:         ps.User.ID,
			Phone:      ps.User.Phone,
			Email:      ps.User.Email,
			Password:   ps.User.Password,
			FirstName:  ps.User.FirstName,
			SecondName: ps.User.SecondName,
			LastName:   ps.User.LastName,
			Avatar:     ps.User.Avatar,
			Role:       ps.User.Role,
			CreatedAt:  ps.User.CreatedAt,
			DeletedAt:  ps.User.DeletedAt,
		},
		Description:    ps.Description,
		Specialization: *spec,
	}
}

func mapSpecializationDomainToDbModel(sp *[]domain.Specialization) *[]specialization {
	var result []specialization
	for _, v := range *sp {
		result = append(result, specialization{ID: v.ID, Name: v.Name})
	}
	return &result
}
