package validators

import (
	"time"

	"github.com/test_server/internal/domain"
)

type meetingRequest struct {
	Id             int64     `json:"id"`
	PsychologistId int64     `json:"psychologist_id" validate:"required,min=1"`
	ClientId       int64     `json:"client_id" validate:"required,min=1"`
	MeetingDate    time.Time `json:"meeting_date"`
	StartTime      float64   `json:"start_time" validate:"required"`
	Endime         float64   `json:"end_time" validate:"required"`
	Status         string    `json:"status" validate:"eq=|eq=not completed|eq=successfully completed|eq=unsuccessfully completed"`
}

func mapMeetingRequestToDomain(meetingRequest *meetingRequest) *domain.Meeting {
	var mtng domain.Meeting
	if meetingRequest.Status == "" {
		meetingRequest.Status = "not completed"
	}

	mtng.PsychologistId = meetingRequest.PsychologistId
	mtng.ClientId = meetingRequest.ClientId
	mtng.MeetingDate = meetingRequest.MeetingDate
	mtng.StartTime = meetingRequest.StartTime
	mtng.EndTime = meetingRequest.Endime
	mtng.Status = meetingRequest.Status
	return &mtng
}
