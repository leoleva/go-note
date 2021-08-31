package repository

import (
	"database/sql"
	"demoproject/src/entity"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// todo: find a better way of mapping, created_at handling..

func (r *UserRepository) Create(user entity.User) (int64, error) {
	h, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	result, err := r.db.Exec("INSERT INTO user (password, email, created_at, uuid) values (?, ?, now(), ?)", string(h), user.Email, user.Uuid)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) GetUserByEmailAndPassword(email string, password string) (entity.User, error) {
	u := entity.User{}

	err := r.db.QueryRow(
		"SELECT id, email, password, uuid FROM user WHERE email = ? LIMIT 1", email,
		).Scan(&u.Id, &u.Email, &u.Password, &u.Uuid)

	if err != nil {
		fmt.Println("SQL error: " +err.Error())

		return entity.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		return entity.User{}, sql.ErrNoRows
	}

	return u, nil
}

func (r *UserRepository) UserExists(email string) bool {
	var id int
	err := r.db.QueryRow("SELECT id FROM user WHERE email = ? LIMIT 1", email).Scan(&id)

	if err != nil {
		if err != sql.ErrNoRows {
			// todo: should be logged

			fmt.Println(err)
		}

		return false
	}

	return true
}

func (r *UserRepository) GetUserByUuid(uuid string) (entity.User, error) {
	u := entity.User{}

	err := r.db.QueryRow(
		"SELECT id, email, uuid FROM user WHERE uuid = ? LIMIT 1", uuid,
	).Scan(&u.Id, &u.Email, &u.Uuid)

	if err != nil {
		return u, err
	}

	return u, nil
}
