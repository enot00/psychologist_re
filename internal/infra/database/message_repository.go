package database

import (
	"github.com/test_server/internal/domain"
	"github.com/upper/db/v4"
	"log"
	"time"
)

const messagesTable = "messages"

type message struct {
	Id       int64     `db:"id,omitempty"`
	ChatId   int64     `db:"chat_id"`
	FromId   int64     `db:"from_id"`
	Text     string    `db:"text"`
	FilePath string    `db:"file_path"`
	Date     time.Time `db:"date"`
}

type MessageRepository interface {
	AddMessage(msg *domain.Message) (*domain.Message, error)
	UpdateMessage(msg *domain.Message) (*domain.Message, error)
	PaginateAll(chatID int64, page uint, pageSize uint) (*[]domain.Message, error)
	DeleteMessage(id int64) error
}

type messageRepository struct {
	sess *db.Session
}

func NewMessageRepository(dbSession *db.Session) MessageRepository {
	return &messageRepository{
		sess: dbSession,
	}
}

func (m *messageRepository) AddMessage(msg *domain.Message) (*domain.Message, error) {
	message := mapChatToMessageDbModel(msg)

	err := (*m.sess).Collection(messagesTable).InsertReturning(message)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return mapMessageDbModelToChat(message), nil
}

func (m *messageRepository) UpdateMessage(msg *domain.Message) (*domain.Message, error) {
	message := mapChatToMessageDbModel(msg)
	err := (*m.sess).Collection(messagesTable).Find(message.Id).Update(message)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return msg, nil
}

func (m *messageRepository) PaginateAll(chatID int64, page uint, pageSize uint) (*[]domain.Message, error) {
	var messages []message

	res := (*m.sess).Collection(messagesTable).
		Find(db.Cond{"chat_id =": chatID}).
		Paginate(pageSize).
		Page(page)

	if err := res.All(&messages); err != nil {
		log.Print("PaginateAll: ", err)
	}

	return mapMessageDbModelToChatCollection(messages), nil
}

func (m *messageRepository) DeleteMessage(id int64) error {
	var msg domain.Message
	res := (*m.sess).Collection(messagesTable).Find(id)
	err := res.One(&msg)
	err = res.Delete()
	return err
}

func mapMessageDbModelToChat(msg *message) *domain.Message {
	return &domain.Message{
		Id:       msg.Id,
		ChatId:   msg.ChatId,
		FromId:   msg.FromId,
		Text:     msg.Text,
		FilePath: msg.FilePath,
		Date:     msg.Date,
	}
}

func mapMessageDbModelToChatCollection(messages []message) *[]domain.Message {
	var result []domain.Message

	for _, msg := range messages {
		newMsg := mapMessageDbModelToChat(&msg)
		result = append(result, *newMsg)
	}
	return &result
}

func mapChatToMessageDbModel(msg *domain.Message) *message {
	return &message{
		Id:       msg.Id,
		ChatId:   msg.ChatId,
		FromId:   msg.FromId,
		Text:     msg.Text,
		FilePath: msg.FilePath,
		Date:     msg.Date,
	}
}
