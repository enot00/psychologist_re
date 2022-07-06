package validators

import (
	"encoding/json"
	"github.com/test_server/internal/domain"
	"github.com/test_server/internal/infra/http/resources"
	"io"
	"log"
)

func MapChatJsonToDto(r io.ReadCloser) (*resources.ChatDto, error) {
	var dto resources.ChatDto

	err := json.NewDecoder(r).Decode(&dto)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &dto, nil
}

func MapChatDtoToModel(dto *resources.ChatDto) *domain.Chat {
	return &domain.Chat{
		CreatorId:    dto.CreatorId,
		Participants: dto.Participants,
	}
}

func MapChatModelToDto(chat *domain.Chat) *resources.ChatDto {
	return &resources.ChatDto{
		Id:           chat.Id,
		CreatorId:    chat.CreatorId,
		Participants: chat.Participants,
	}
}
