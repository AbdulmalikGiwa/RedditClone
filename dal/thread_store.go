package dal

import (
	"fmt"

	redditclone "github.com/AbdulmalikGiwa/RedditClone"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

//NewThreadStore returns a new ThreadStore(Obvious ikr, golint keeps forcing me to comment)
func NewThreadStore(db *sqlx.DB) ThreadStore {
	return ThreadStore{
		DB: db,
	}
}

// ThreadStore extends methods and types of sqlx.DB
type ThreadStore struct {
	*sqlx.DB
}

//Thread selects row of "id" from threads table and embeds in struct
func (s *ThreadStore) Thread(id uuid.UUID) (redditclone.Thread, error) {
	var t redditclone.Thread
	err := s.Get(&t, `SELECT * FROM threads WHERE id= $1`, id)
	if err != nil {
		return redditclone.Thread{}, fmt.Errorf("error getting thread: %w", err)
	}
	return t, nil
}

//Threads selects all rows from threads table and embeds in slice
func (s *ThreadStore) Threads() ([]redditclone.Thread, error) {
	var t []redditclone.Thread
	err := s.Select(&t, `SELECT * FROM threads`)
	if err != nil {
		return []redditclone.Thread{}, fmt.Errorf("error getting thread: %w", err)
	}
	return t, nil
}

//CreateThread creates new thread in the database
func (s *ThreadStore) CreateThread(t *redditclone.Thread) error {
	err := s.Get(t, `INSERT INTO threads VALUES ($1,$2,$3) RETURNING *`,
		t.ID,
		t.Title,
		t.Description)
	if err != nil {
		return fmt.Errorf("Error creating thread: %w", err)
	}
	return nil
}

//UpdateThread updates existing thread in db
func (s *ThreadStore) UpdateThread(t *redditclone.Thread) error {
	err := s.Get(t, `UPDATE threads SET title = $1, description = $2 WHERE id = $3 RETURNING *`,
		t.Title,
		t.Description,
		t.ID)
	if err != nil {
		return fmt.Errorf("Error updating thread: %w", err)
	}
	return nil
}

//DeleteThread deletes existing thread from db
func (s *ThreadStore) DeleteThread(id uuid.UUID) error {
	_, err := s.Exec(`DELETE FROM threads WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("Error Deleting thread: %w ", err)
	}
	return nil
}
