package database

import (
	"github.com/lib/pq"
	"github.com/test_server/internal/domain"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"sort"
)

type ChatRepository interface {
	Add(chat *domain.Chat) (*domain.Chat, error)
	FindChat(id uint64) (*domain.Chat, error)
	FindUserChats(id uint64) ([]domain.Chat, error)
	Delete(id uint64) error
}

const chatsTable = "chats"

type result struct {
	Id           uint64                `db:"id,omitempty"`
	CreatorId    uint64                `db:"user_id"`
	Participants postgresql.Int64Array `db:"participants"`
}

type insert struct {
	Id           uint64      `db:"id,omitempty"`
	CreatorId    uint64      `db:"user_id"`
	Participants interface{} `db:"participants"`
}

type postgresChatRepository struct {
	db *db.Session
}

func NewChatRepository(db *db.Session) ChatRepository {
	return &postgresChatRepository{db}
}

func (r *postgresChatRepository) Add(c *domain.Chat) (*domain.Chat, error) {
	insert := r.mapModelToRow(c)

	err := (*r.db).Collection(chatsTable).InsertReturning(insert)
	if err != nil {
		log.Printf("postgresChatRepository.Add(): %s", err.Error())
		return nil, err
	}

	return r.FindChat(insert.Id)
}

func (r *postgresChatRepository) FindChat(id uint64) (*domain.Chat, error) {
	var row result

	err := (*r.db).Collection(chatsTable).
		Find(db.Cond{"id": id}).
		One(&row)
	if err != nil {
		log.Printf("postgresChatRepository.FindChat(): %s", err.Error())

		return nil, err
	}

	return r.mapRowToModel(&row), nil
}

func (r *postgresChatRepository) FindUserChats(id uint64) ([]domain.Chat, error) {
	var rows []result

	err := (*r.db).Collection(chatsTable).
		Find("? = ANY(participants)", id).
		All(&rows)

	if err != nil {
		log.Printf("postgresChatRepository.FindUserChats(): %s", err.Error())

		return nil, err
	}

	return r.mapRowsToCollection(rows), nil
}

func (r *postgresChatRepository) Delete(id uint64) error {
	err := (*r.db).Collection(chatsTable).
		Find(db.Cond{"id": id}).
		Delete()
	if err != nil {
		log.Printf("postgresChatRepository.Delete(): %s", err.Error())
		return err
	}

	return nil
}

func (r *postgresChatRepository) mapModelToRow(c *domain.Chat) *insert {
	sorted := c.Participants
	sort.SliceStable(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	return &insert{
		Id:           c.Id,
		CreatorId:    c.CreatorId,
		Participants: pq.Array(sorted),
	}
}

func (r *postgresChatRepository) mapRowToModel(row *result) *domain.Chat {
	var p []uint64

	for _, v := range row.Participants {
		p = append(p, uint64(v))
	}

	return &domain.Chat{
		Id:           row.Id,
		CreatorId:    row.CreatorId,
		Participants: p,
	}
}

func (r *postgresChatRepository) mapRowsToCollection(rows []result) []domain.Chat {
	var chats []domain.Chat

	for _, row := range rows {
		chats = append(chats, *r.mapRowToModel(&row))
	}

	return chats
}
