package domain

import "time"

type Message struct {
	Id       int64
	ChatId   int64
	FromId   int64
	Text     string
	FilePath string
	Date     time.Time
}
