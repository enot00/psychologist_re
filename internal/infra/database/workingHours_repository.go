package database

import (
	"github.com/test_server/internal/domain"
	"github.com/upper/db/v4"
	"time"
)

type workingHours struct {
	ID             int64         `db:"id,omitempty"`
	PsychologistID int64         `db:"psychologist_id"`
	WeekDay        *time.Weekday `db:"week_day,omitempty"`
	Date           *time.Time    `db:"date,omitempty"`
	StartTime      time.Time     `db:"start_time"`
	EndTime        time.Time     `db:"end_time"`
}

type WorkingHoursRepository interface {
	GetOneByID(id int64) (*[]domain.WorkingHours, error)
	Save(workingHours *[]domain.WorkingHours) (*[]domain.WorkingHours, error)
	Update(id int64, workingHours *[]domain.WorkingHours) error
	Delete(id int64) error
}
type workingHoursRepository struct {
	coll   db.Session
	whColl db.Collection
}

func NewWorkingHoursRepository(dbSession *db.Session) WorkingHoursRepository {
	return &workingHoursRepository{
		coll:   *dbSession,
		whColl: (*dbSession).Collection("working_hours"),
	}
}
func (r *workingHoursRepository) GetOneByID(id int64) (*[]domain.WorkingHours, error) {

	var wh []workingHours
	err := r.whColl.Find(db.Cond{"psychologist_id": id}).All(&wh)
	if err != nil {
		return nil, err
	}
	return mapWorkingHoursDbModelToDomain(&wh), nil
}
func (r *workingHoursRepository) Save(wh *[]domain.WorkingHours) (*[]domain.WorkingHours, error) {
	whWithID := mapWorkingHoursDomainToDbModel(wh)
	err := r.whColl.InsertReturning(whWithID)
	whDomain := mapWorkingHoursDbModelToDomain(whWithID)
	if err != nil {
		return nil, err
	}
	return whDomain, nil
}
func (r *workingHoursRepository) Update(id int64, hours *[]domain.WorkingHours) error {
	wh := mapWorkingHoursDomainToDbModel(hours)
	err := r.Delete(id)
	if err != nil {
		return err
	}
	for _, v := range *wh {
		ps := workingHours{
			PsychologistID: id,
			WeekDay:        v.WeekDay,
			Date:           v.Date,
			StartTime:      v.StartTime,
			EndTime:        v.EndTime,
		}

		err = r.whColl.InsertReturning(&ps)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *workingHoursRepository) Delete(id int64) error {

	err := r.whColl.Find(db.Cond{"psychologist_id": id}).Delete()
	if err != nil {
		return err
	}
	return nil
}

func mapWorkingHoursDbModelToDomain(wh *[]workingHours) *[]domain.WorkingHours {
	var result []domain.WorkingHours

	for _, v := range *wh {
		result = append(result, domain.WorkingHours{
			ID:        v.ID,
			UserID:    v.PsychologistID,
			WeekDay:   v.WeekDay,
			Date:      v.Date,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return &result
}
func mapWorkingHoursDomainToDbModel(wh *[]domain.WorkingHours) *[]workingHours {
	var result []workingHours
	for _, v := range *wh {
		result = append(result, workingHours{
			ID:             v.ID,
			PsychologistID: v.UserID,
			WeekDay:        v.WeekDay,
			Date:           v.Date,
			StartTime:      v.StartTime,
			EndTime:        v.EndTime,
		})
	}
	return &result
}
