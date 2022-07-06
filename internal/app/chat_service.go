package app

import (
	"github.com/test_server/internal/infra/database"
	"github.com/test_server/internal/infra/http/resources"
	mapper "github.com/test_server/internal/infra/http/validators"
)

type ChatService struct {
	r *database.ChatRepository
}

func NewChatService(r *database.ChatRepository) *ChatService {
	return &ChatService{r}
}

func (s *ChatService) Add(dto *resources.ChatDto) (*resources.ChatDto, error) {
	chat, err := (*s.r).Add(mapper.MapChatDtoToModel(dto))
	if err != nil {
		return nil, err
	}

	return mapper.MapChatModelToDto(chat), nil
}

func (s *ChatService) FindChat(id uint64) (*resources.ChatDto, error) {
	c, err := (*s.r).FindChat(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapChatModelToDto(c), nil
}

func (s *ChatService) FindUserChats(id uint64) ([]resources.ChatDto, error) {
	var chatsDto []resources.ChatDto
	chats, err := (*s.r).FindUserChats(id)
	if err != nil {
		return nil, err
	}

	for _, v := range chats {
		chatsDto = append(chatsDto, *mapper.MapChatModelToDto(&v))
	}

	return chatsDto, nil
}

func (s *ChatService) Delete(id uint64) error {
	return (*s.r).Delete(id)
}
