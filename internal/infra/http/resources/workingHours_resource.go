package resources

import (
	"time"
)

type WorkingHoursDTO struct {
	ID        int64         `json:"id"`
	WeekDay   *time.Weekday `json:"week_day"`
	Date      *time.Time    `json:"date"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
}
