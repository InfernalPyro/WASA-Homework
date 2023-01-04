package database

// Make user "id" follow user "followId"
func (db *appdbimpl) GetComment(commentId uint64, photoId uint64) (Comment, error) {

	var comment Comment

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return comment, err
	}

	// Query to remove the like
	row, err := tx.Query("SELECT * FROM comment where commentId = ? and photoId = ?", commentId, photoId)
	if err != nil {
		if tx.Rollback() != nil {
			return comment, err
		}
		return comment, err
	}
	defer func() { _ = row.Close() }()
	if row.Next() {
		err = row.Scan(&comment.CommentId, &comment.UserId, &comment.PhotoId, &comment.Content, &comment.Time)
		if err != nil {
			return comment, err
		}
	}

	if err = row.Err(); err != nil {
		return comment, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return comment, err
	}

	return comment, nil
}
