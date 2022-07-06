package validators

import (
	"github.com/test_server/internal/domain"
	"time"
)

type messageRequest struct {
	Id       int64     `json:"id"`
	ChatId   int64     `json:"chat_id" validate:"required"`
	FromId   int64     `json:"from_id" validate:"required"`
	Text     string    `json:"text" validate:"required"`
	FilePath string    `json:"file"`
	Date     time.Time `json:"date" validate:"required"`
}

func mapMessageRequestToChat(messageRequest *messageRequest) *domain.Message {
	var msg domain.Message
	msg.Id = messageRequest.Id
	msg.ChatId = messageRequest.ChatId
	msg.FromId = messageRequest.FromId
	msg.Text = messageRequest.Text
	msg.FilePath = messageRequest.FilePath
	msg.Date = messageRequest.Date

	return &msg
}
