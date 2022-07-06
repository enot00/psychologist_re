package database

import (
	"errors"
	"log"
	"time"

	"github.com/test_server/internal/domain"
	"github.com/upper/db/v4"
)

type meeting struct {
	Id             int64     `db:"id,omitempty"`
	PsychologistId int64     `db:"psychologist_id"`
	ClientId       int64     `db:"client_id"`
	MetingDate     time.Time `db:"meeting_date"`
	StartTime      float64   `db:"start_time"`
	EndTime        float64   `db:"end_time"`
	Status         string    `db:"status"`
}

type MeetingRepository interface {
	SaveMeetingByPsychologist(meeting *domain.Meeting) (*domain.Meeting, error)
	SaveMeetingByClient(meeting *domain.Meeting) (*domain.Meeting, error)
	PaginateAllMeetings(page uint, pageSize uint) (*[]domain.Meeting, error)
	PaginateAllPsychologistMeetings(psychologist_id int64, page uint, pageSize uint) (*[]domain.Meeting, error)
	PaginateAllClientMeetings(clientId int64, page uint, pageSize uint) (*[]domain.Meeting, error)
	FindOneMeeting(meetingId int64) (*domain.Meeting, error)
	UpdateMeetingByPsychologist(meeting *domain.Meeting) error
	UpdateMeetingByClient(meeting *domain.Meeting) error
	DeleteMeetingByPsychologist(meetingId, psychologistId int64) error
	DeleteMeetingByClient(meetingId, clientId int64) error
}

type meetingRepository struct {
	meetingColl      db.Collection
	psychologistColl db.Collection
	clientColl       db.Collection
	workingHoursColl db.Collection
}

func NewMeetingRepository(dbSession *db.Session) MeetingRepository {
	return &meetingRepository{
		meetingColl:      (*dbSession).Collection("meeting"),
		psychologistColl: (*dbSession).Collection("psychologist"),
		clientColl:       (*dbSession).Collection("client"),
		workingHoursColl: (*dbSession).Collection("working_hours"),
	}
}

func (r *meetingRepository) SaveMeetingByPsychologist(meeting *domain.Meeting) (*domain.Meeting, error) {
	// check 1: client exists
	clientExists, _ := r.clientColl.Find(db.Cond{"id": meeting.ClientId}).Exists()
	if !clientExists {
		err := errors.New("the client with this id does not exist")
		log.Print(err)
		return nil, err
	}

	// check 2: psychologist's workinghours exists
	weekDay := int(meeting.MeetingDate.Weekday())
	if weekDay == 0 { // if Sunday
		weekDay = 7
	}

	res := r.workingHoursColl.Find()
	psychologistWorks, _ := res.And(
		"psychologist_id = ? AND (date = ? OR (week_days = ? AND date IS NULL)) AND (start_time <= ? AND end_time >= ?)",
		meeting.PsychologistId, meeting.MeetingDate, weekDay, meeting.StartTime, meeting.EndTime,
	).Exists()
	if !psychologistWorks {
		err := errors.New("the psychologist does not work on this day, at this time")
		log.Print(err)
		return nil, err
	}

	// check 3: does the meeting overlap with an existing meeting
	meetingOverlaps, _ := r.meetingColl.Find().And(
		"(psychologist_id = ? OR client_id = ?) AND meeting_date = ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?))",
		meeting.PsychologistId, meeting.ClientId, meeting.MeetingDate, meeting.StartTime, meeting.StartTime, meeting.EndTime, meeting.EndTime,
	).Exists()
	if meetingOverlaps {
		err := errors.New("psychologist or client already has a meeting on this day, in this time range")
		log.Print(err)
		return nil, err
	}

	// if all checks passed, create the meeting
	mtng := mapDomainToMeetingDbModel(meeting)
	err := r.meetingColl.InsertReturning(mtng)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapMeetingDbModelToDomain(mtng), nil
}

func (r *meetingRepository) SaveMeetingByClient(meeting *domain.Meeting) (*domain.Meeting, error) {
	// check 1: psychologist exists
	psychologistExists, _ := r.psychologistColl.Find(db.Cond{"id": meeting.PsychologistId}).Exists()
	if !psychologistExists {
		err := errors.New("the psychologist with this id does not exist")
		log.Print(err)
		return nil, err
	}

	// check 2: psychologist's workinghours exists
	weekDay := int(meeting.MeetingDate.Weekday())
	if weekDay == 0 { // if Sunday
		weekDay = 7
	}

	res := r.workingHoursColl.Find()
	psychologistWorks, _ := res.And(
		"psychologist_id = ? AND (date = ? OR (week_days = ? AND date IS NULL)) AND (start_time <= ? AND end_time >= ?)",
		meeting.PsychologistId, meeting.MeetingDate, weekDay, meeting.StartTime, meeting.EndTime,
	).Exists()
	if !psychologistWorks {
		err := errors.New("the psychologist does not work on this day, at this time")
		log.Print(err)
		return nil, err
	}

	// check 3: does the meeting overlap with an existing meeting
	meetingOverlaps, _ := r.meetingColl.Find().And(
		"(psychologist_id = ? OR client_id = ?) AND meeting_date = ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?))",
		meeting.PsychologistId, meeting.ClientId, meeting.MeetingDate, meeting.StartTime, meeting.StartTime, meeting.EndTime, meeting.EndTime,
	).Exists()
	if meetingOverlaps {
		err := errors.New("psychologist or client already has a meeting on this day, in this time range")
		log.Print(err)
		return nil, err
	}

	// if all checks passed, create the meeting
	mtng := mapDomainToMeetingDbModel(meeting)
	err := r.meetingColl.InsertReturning(mtng)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapMeetingDbModelToDomain(mtng), nil
}

func (r *meetingRepository) PaginateAllMeetings(page uint, pageSize uint) (*[]domain.Meeting, error) {
	var meetings []meeting
	err := r.meetingColl.Find().Paginate(pageSize).Page(page).
		OrderBy("status", "meeting_date", "start_time").All(&meetings)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapMeetingDbModelToDomainCollection(&meetings), nil
}

func (r *meetingRepository) PaginateAllPsychologistMeetings(psychologistId int64, page uint, pageSize uint) (*[]domain.Meeting, error) {
	// check psychologist exists
	psychologistExists, _ := r.psychologistColl.Find(db.Cond{"id": psychologistId}).Exists()
	if !psychologistExists {
		err := errors.New("the psychologist with this id does not exist")
		log.Print(err)
		return nil, err
	}

	var meetings []meeting
	err := r.meetingColl.Find(db.Cond{"psychologist_id": psychologistId}).
		Paginate(pageSize).Page(page).OrderBy("status").All(&meetings)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapMeetingDbModelToDomainCollection(&meetings), nil
}

func (r *meetingRepository) PaginateAllClientMeetings(clientId int64, page uint, pageSize uint) (*[]domain.Meeting, error) {
	//check client exists
	clientExists, _ := r.clientColl.Find(db.Cond{"id": clientId}).Exists()
	if !clientExists {
		err := errors.New("the client with this id does not exist")
		log.Print(err)
		return nil, err
	}

	var meetings []meeting
	err := r.meetingColl.Find(db.Cond{"client_id": clientId}).
		Paginate(pageSize).Page(page).OrderBy("status").All(&meetings)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapMeetingDbModelToDomainCollection(&meetings), nil
}

func (r *meetingRepository) FindOneMeeting(meetingId int64) (*domain.Meeting, error) {
	var mtng *meeting
	err := r.meetingColl.Find(db.Cond{"id": meetingId}).One(&mtng)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapMeetingDbModelToDomain(mtng), nil
}

func (r *meetingRepository) UpdateMeetingByPsychologist(meeting *domain.Meeting) error {
	// check 1: meeting exists and the psychologist has the permission to edit this meeting
	meetingExists, _ := r.meetingColl.Find(
		db.Cond{"id": meeting.Id, "psychologist_id": meeting.PsychologistId},
	).Exists()
	if !meetingExists {
		err := errors.New("the meeting with this id does not exist or psychologist has no permission to edit this meeting")
		log.Print(err)
		return err
	}

	// check 2: client exists
	clientExists, _ := r.clientColl.Find(db.Cond{"id": meeting.ClientId}).Exists()
	if !clientExists {
		err := errors.New("the client with this id does not exist")
		log.Print(err)
		return err
	}

	// check 3: psychologist's workinghours exists
	weekDay := int(meeting.MeetingDate.Weekday())
	if weekDay == 0 { // if Sunday
		weekDay = 7
	}

	res := r.workingHoursColl.Find()
	psychologistWorks, _ := res.And(
		"psychologist_id = ? AND (date = ? OR (week_days = ? AND date IS NULL)) AND (start_time <= ? AND end_time >= ?)",
		meeting.PsychologistId, meeting.MeetingDate, weekDay, meeting.StartTime, meeting.EndTime,
	).Exists()
	if !psychologistWorks {
		err := errors.New("the psychologist does not work on this day, at this time")
		log.Print(err)
		return err
	}

	// check 4: does the meeting overlap with an existing meeting
	meetingOverlaps, _ := r.meetingColl.Find().And(
		"id != ? AND (psychologist_id = ? OR client_id = ?) AND meeting_date = ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?))",
		meeting.Id, meeting.PsychologistId, meeting.ClientId, meeting.MeetingDate, meeting.StartTime, meeting.StartTime, meeting.EndTime, meeting.EndTime,
	).Exists()
	if meetingOverlaps {
		err := errors.New("psychologist or client already has another meeting on this day, in this time range")
		log.Print(err)
		return err
	}

	// if all checks passed, update the meeting
	updtMtng := mapDomainToMeetingDbModel(meeting)
	err := r.meetingColl.UpdateReturning(updtMtng)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (r *meetingRepository) UpdateMeetingByClient(meeting *domain.Meeting) error {
	// check 1: meeting exists and the client has the permission to edit this meeting
	meetingExists, _ := r.meetingColl.Find(
		db.Cond{"id": meeting.Id, "client_id": meeting.ClientId},
	).Exists()
	if !meetingExists {
		err := errors.New("the meeting with this id does not exist or client has no permission to edit this meeting")
		log.Print(err)
		return err
	}

	// check 2: client exists
	psychologistExists, _ := r.psychologistColl.Find(db.Cond{"id": meeting.PsychologistId}).Exists()
	if !psychologistExists {
		err := errors.New("the psychologist with this id does not exist")
		log.Print(err)
		return err
	}

	// check 3: psychologist's workinghours exists
	weekDay := int(meeting.MeetingDate.Weekday())
	if weekDay == 0 { // if Sunday
		weekDay = 7
	}

	res := r.workingHoursColl.Find()
	psychologistWorks, _ := res.And(
		"psychologist_id = ? AND (date = ? OR (week_days = ? AND date IS NULL)) AND (start_time <= ? AND end_time >= ?)",
		meeting.PsychologistId, meeting.MeetingDate, weekDay, meeting.StartTime, meeting.EndTime,
	).Exists()
	if !psychologistWorks {
		err := errors.New("the psychologist does not work on this day, at this time")
		log.Print(err)
		return err
	}

	// check 4: does the meeting overlap with an existing meeting
	meetingOverlaps, _ := r.meetingColl.Find().And(
		"id != ? AND (psychologist_id = ? OR client_id = ?) AND meeting_date = ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?))",
		meeting.Id, meeting.PsychologistId, meeting.ClientId, meeting.MeetingDate, meeting.StartTime, meeting.StartTime, meeting.EndTime, meeting.EndTime,
	).Exists()
	if meetingOverlaps {
		err := errors.New("psychologist or client already has another meeting on this day, in this time range")
		log.Print(err)
		return err
	}

	// if all checks passed, update the meeting
	updtMtng := mapDomainToMeetingDbModel(meeting)
	err := r.meetingColl.UpdateReturning(updtMtng)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (r *meetingRepository) DeleteMeetingByPsychologist(meetingId, psychologistId int64) error {
	// check meeting exists and the psychologist has the permission to delete this meeting
	res := r.meetingColl.Find(db.Cond{"id": meetingId, "psychologist_id": psychologistId})
	meetingExists, _ := res.Exists()
	if !meetingExists {
		err := errors.New("the meeting with this id does not exist or psychologist has no permission to delete this meeting")
		log.Print(err)
		return err
	}

	err := res.Delete()
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (r *meetingRepository) DeleteMeetingByClient(meetingId, clientId int64) error {
	// check meeting exists and the client has the permission to delete this meeting
	res := r.meetingColl.Find(db.Cond{"id": meetingId, "client_id": clientId})
	meetingExists, _ := res.Exists()
	if !meetingExists {
		err := errors.New("the meeting with this id does not exist or client has no permission to delete this meeting")
		log.Print(err)
		return err
	}

	err := res.Delete()
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func mapMeetingDbModelToDomain(meeting *meeting) *domain.Meeting {
	return &domain.Meeting{
		Id:             meeting.Id,
		PsychologistId: meeting.PsychologistId,
		ClientId:       meeting.ClientId,
		MeetingDate:    meeting.MetingDate,
		StartTime:      meeting.StartTime,
		EndTime:        meeting.EndTime,
		Status:         meeting.Status,
	}
}

func mapMeetingDbModelToDomainCollection(meetings *[]meeting) *[]domain.Meeting {
	var result []domain.Meeting

	for _, mtng := range *meetings {
		newMtng := mapMeetingDbModelToDomain(&mtng)
		result = append(result, *newMtng)
	}
	return &result
}

func mapDomainToMeetingDbModel(mtng *domain.Meeting) *meeting {
	return &meeting{
		Id:             mtng.Id,
		PsychologistId: mtng.PsychologistId,
		ClientId:       mtng.ClientId,
		MetingDate:     mtng.MeetingDate,
		StartTime:      mtng.StartTime,
		EndTime:        mtng.EndTime,
		Status:         mtng.Status,
	}
}
