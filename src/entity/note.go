package entity

import (
	"time"
)

type Note struct {
	Id    int64
	Title string
	Text  string
	UserId int64
	CreatedAt time.Time
	Uuid string
}

func NewNote(title string, text string, userId int64, createdAt time.Time, uuid string) *Note {
	return &Note {
		Title:     title,
		Text:      text,
		UserId:      userId,
		CreatedAt: createdAt,
		Uuid: uuid,
	}
}

func (n *Note) GetId() int64 {
	return n.Id
}

func (n *Note) GetTitle() string {
	return n.Title
}

func (n *Note) GetText() string {
	return n.Text
}

func (n *Note) GetUserId() int64 {
	return n.UserId
}

func (n *Note) GetCreatedAt() time.Time {
	return n.CreatedAt
}

func (n *Note) GetUuid() string {
	return n.Uuid
}

func (n *Note) UpdateId(id int64) *Note {
	n.Id = id

	return n
}

func (n *Note) ToMap() map[string]interface{} {
	noteMap := make(map[string]interface{})

	noteMap["title"] = n.Title
	noteMap["text"] = n.Text
	noteMap["createdAt"] = n.CreatedAt.Format("20060102150405")

	return noteMap
}
