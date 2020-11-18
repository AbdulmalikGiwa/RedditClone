package postgres

import "github.com/jmoiron/sqlx"

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
