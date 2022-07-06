package domain

import (
	"time"
)

type WorkingHours struct {
	ID        int64
	UserID    int64
	WeekDay   *time.Weekday
	Date      *time.Time
	StartTime time.Time
	EndTime   time.Time
}
