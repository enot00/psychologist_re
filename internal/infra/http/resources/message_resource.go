package resources

import (
	"github.com/test_server/internal/domain"
	"time"
)

type MessageDto struct {
	Id       int64     `json:"id,omitempty"`
	ChatId   int64     `json:"chat_id"`
	FromId   int64     `json:"from_id"`
	Text     string    `json:"text"`
	FilePath string    `json:"file"`
	Date     time.Time `json:"date"`
}

func MapChatToMessageDto(msg *domain.Message) *MessageDto {
	return &MessageDto{
		Id:       msg.Id,
		ChatId:   msg.ChatId,
		FromId:   msg.FromId,
		Text:     msg.Text,
		FilePath: msg.FilePath,
		Date:     msg.Date,
	}
}

func MapChatToMessageDtoCollection(msg *[]domain.Message) *[]MessageDto {
	var result []MessageDto
	for _, t := range *msg {
		dto := MapChatToMessageDto(&t)
		result = append(result, *dto)
	}
	return &result
}
