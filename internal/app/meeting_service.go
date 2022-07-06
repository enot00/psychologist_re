package app

import (
	"log"

	"github.com/test_server/internal/domain"
	"github.com/test_server/internal/infra/database"
)

type MeetingService interface {
	SaveMeetingByPsychologist(meeting *domain.Meeting) (*domain.Meeting, error)
	SaveMeetingByClient(meeting *domain.Meeting) (*domain.Meeting, error)
	PaginateAllMeetings(page uint, pageSize uint) (*[]domain.Meeting, error)
	PaginateAllPsychologistMeetings(psychologistId int64, page uint, pageSize uint) (*[]domain.Meeting, error)
	PaginateAllClientMeetings(clientId int64, page uint, pageSize uint) (*[]domain.Meeting, error)
	FindOneMeeting(meetingId int64) (*domain.Meeting, error)
	UpdateMeetingByPsychologist(meeting *domain.Meeting) error
	UpdateMeetingByClient(meeting *domain.Meeting) error
	DeleteMeetingByPsychologist(meetingId, psychologistId int64) error
	DeleteMeetingByClient(meetingId, clientId int64) error
}

type meetingService struct {
	meetingRepo *database.MeetingRepository
}

func NewMeetingService(m *database.MeetingRepository) MeetingService {
	return &meetingService{
		meetingRepo: m,
	}
}

func (s *meetingService) SaveMeetingByPsychologist(meeting *domain.Meeting) (*domain.Meeting, error) {
	meeting, err := (*s.meetingRepo).SaveMeetingByPsychologist(meeting)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return meeting, nil
}

func (s *meetingService) SaveMeetingByClient(meeting *domain.Meeting) (*domain.Meeting, error) {
	meeting, err := (*s.meetingRepo).SaveMeetingByClient(meeting)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return meeting, nil
}

func (s *meetingService) PaginateAllMeetings(page uint, pageSize uint) (*[]domain.Meeting, error) {
	meetings, err := (*s.meetingRepo).PaginateAllMeetings(page, pageSize)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return meetings, nil
}

func (s *meetingService) PaginateAllPsychologistMeetings(psychologistId int64, page uint, pageSize uint) (*[]domain.Meeting, error) {
	meetings, err := (*s.meetingRepo).PaginateAllPsychologistMeetings(
		psychologistId, page, pageSize,
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return meetings, nil
}

func (s *meetingService) PaginateAllClientMeetings(clientId int64, page uint, pageSize uint) (*[]domain.Meeting, error) {
	meetings, err := (*s.meetingRepo).PaginateAllClientMeetings(
		clientId, page, pageSize,
	)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return meetings, nil
}

func (s *meetingService) FindOneMeeting(meetingId int64) (*domain.Meeting, error) {
	meeting, err := (*s.meetingRepo).FindOneMeeting(meetingId)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return meeting, nil
}

func (s *meetingService) UpdateMeetingByPsychologist(meeting *domain.Meeting) error {
	err := (*s.meetingRepo).UpdateMeetingByPsychologist(meeting)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (s *meetingService) UpdateMeetingByClient(meeting *domain.Meeting) error {
	err := (*s.meetingRepo).UpdateMeetingByClient(meeting)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (s *meetingService) DeleteMeetingByPsychologist(meetingId, psychologistId int64) error {
	err := (*s.meetingRepo).DeleteMeetingByPsychologist(meetingId, psychologistId)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (s *meetingService) DeleteMeetingByClient(meetingId, clientId int64) error {
	err := (*s.meetingRepo).DeleteMeetingByClient(meetingId, clientId)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
