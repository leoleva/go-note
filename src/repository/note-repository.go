package repository

import (
	"database/sql"
	"demoproject/src/entity"
	"fmt"
	"time"
)

type NoteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) *NoteRepository {
	return &NoteRepository{
		db: db,
	}
}

func (r *NoteRepository) Create(note entity.Note) (entity.Note, error) {
	exec, err := r.db.Exec(
		"INSERT INTO note (title, text, user_id, created_at, uuid) values (?, ?, ?, ?, ?)",
		note.GetTitle(),
		note.GetText(),
		note.GetUserId(),
		note.GetCreatedAt(),
		note.GetUuid(),
	)

	if err != nil {
		return entity.Note{}, err
	}

	id, err := exec.LastInsertId()

	note.UpdateId(id)

	return note, nil
}

func (r *NoteRepository) NoteExistsByUserIdAndUuid(userId int64, uuid string) bool {
	var id int

	err := r.db.QueryRow("SELECT id FROM note WHERE user_id = ? AND uuid = ?", userId, uuid).Scan(&id)

	if err != nil {
		if err != sql.ErrNoRows {
			// todo: should be logged

			fmt.Println(err)
		}

		return false
	}

	return true
}

func (r *NoteRepository) DeleteByUserIdAndUuid(userId int64, uuid string) error {
	_, err := r.db.Exec("DELETE FROM note WHERE user_id = ? AND uuid = ?", userId, uuid)

	return err
}

func (r *NoteRepository) GetNoteByUserIdAndUuid(userId int64, uuid string) (note entity.Note, err error) {
	var createdAt string

	err = r.db.QueryRow(
		"SELECT id, title, text, uuid, user_id, created_at FROM note WHERE user_id = ? AND uuid = ?",
		userId,
		uuid,
		).Scan(
			&note.Id,
			&note.Title,
			&note.Text,
			&note.Uuid,
			&note.UserId,
			&createdAt,
			)

	if err != nil {
		return note, err
	}

	createdAtTime, err := time.Parse("2006-01-02 15:04:05", createdAt)

	if err != nil {
		return note, err
	}

	note.CreatedAt = createdAtTime

	return note, nil
}
