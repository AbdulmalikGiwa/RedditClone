package dal

import (
	"fmt"

	redditclone "github.com/AbdulmalikGiwa/RedditClone"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

//NewPostStore returns PostStore
func NewPostStore(db *sqlx.DB) PostStore {
	return PostStore{
		DB: db,
	}
}

//PostStore type extends sqlx.DB and interacts with Posts Table
type PostStore struct {
	*sqlx.DB
}

//Post returns new post
func (s *PostStore) Post(id uuid.UUID) (redditclone.Post, error) {
	var p redditclone.Post
	err := s.Get(&p, `SELECT $1 FROM posts`, id)
	if err != nil {
		return redditclone.Post{}, fmt.Errorf("Error getting post: %w", err)
	}
	return p, nil
}

//PostsbyThread gets all posts in a thread
func (s *PostStore) PostsbyThread(threadID uuid.UUID) ([]redditclone.Post, error) {
	var p []redditclone.Post
	err := s.Select(&p, `SELECT * FROM posts WHERE thread_id = $1`, threadID)
	if err != nil {
		return []redditclone.Post{}, fmt.Errorf("Error fetching posts: %w", err)
	}
	return p, nil
}

//CreatePost creates a new post
func (s *PostStore) CreatePost(t *redditclone.Post) error {
	err := s.Get(t, `INSERT INTO posts VALUES ($1,$2,$3,$4,$5) RETURNING *`,
		t.ID,
		t.ThreadID,
		t.Title,
		t.Content,
		t.Votes)
	if err != nil {
		return fmt.Errorf("Error creating post: %w", err)
	}
	return nil
}

//UpdatePost updates posts
func (s *PostStore) UpdatePost(t *redditclone.Post) error {
	err := s.Get(t, `UPDATE posts SET thread_id=$1, title=$2, content=$3, votes=$4 WHERE id=$5`,
		t.ThreadID,
		t.Title,
		t.Content,
		t.Votes,
		t.ID)
	if err != nil {
		return fmt.Errorf("Error updating post: %w", err)
	}
	return nil
}

//DeletePost deletes posts
func (s *PostStore) DeletePost(id uuid.UUID) error {
	_, err := s.Exec(`DELETE FROM posts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("Error deleting post: %w", err)
	}
	return nil
}
