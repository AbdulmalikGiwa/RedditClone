package dal

import (
	"fmt"
	redditclone "github.com/AbdulmalikGiwa/RedditClone"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

//NewCommentStore returns CommentStore
func NewCommentStore(db *sqlx.DB) CommentStore {
	return CommentStore{
		DB: db,
	}
}

//CommentStore extends sqlx.DB
type CommentStore struct {
	*sqlx.DB
}

func (c *CommentStore) Comment(id uuid.UUID) (redditclone.Comment, error) {
	var d redditclone.Comment
	err := c.Get(&d, `SELECT * FROM comments WHERE id=$1`, id)
	if err != nil {
		return redditclone.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}
	return  d, nil
}

func (c *CommentStore) CommentsbyPost(postID uuid.UUID) ([]redditclone.Comment, error) {
	var d []redditclone.Comment
	err := c.Select(&d, `SELECT * FROM TABLE`)
	if err != nil {
		return []redditclone.Comment{}, fmt.Errorf("error getting all comments: %w", err)
	}
	return d, nil
}

func (c *CommentStore) CreateComment(t *redditclone.Comment) error {
	err := c.Get(t, `INSERT INTO  comments VALUES ($1,$2,$3,$4) RETURNING * `,
		t.ID,
		t.PostID,
		t.Content,
		t.Votes)
	if err != nil {
		return fmt.Errorf("error creating comment: %w", err)
	}
	return nil
}

func (c CommentStore) UpdateComment(t *redditclone.Comment) error {
	err := c.Get(t, `UPDATE comments SET post_id=$1, content=$2, votes=$3 WHERE id=$4 RETURNING *`,
		t.PostID,
		t.Content,
		t.Votes,
		t.ID)
	if err != nil {
		return fmt.Errorf("error updating comment: %w", err)
	}
	return nil
}

func (c *CommentStore) DeleteComment(id uuid.UUID) error {
	_, err := c.Exec(`DELETE FROM comments WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("Error Deleting comment: %w ", err)
	}
	return nil

}

