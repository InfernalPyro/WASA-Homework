package database

// Make user "id" follow user "followId"
func (db *appdbimpl) CommentPhoto(photoId uint64, comment Comment) (Comment, error) {

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return comment, err
	}

	// Query to insert the new like
	_, err = tx.Exec("Insert into comment (photoId, userId, content, time) values (?,?,?,?)", photoId, comment.UserId, comment.Content, comment.Time)
	if err != nil {
		if err.Error() == ErrForeignKey {
			err = ErrPhotoNotFound
		}
		if tx.Rollback() != nil {
			return comment, err
		}
		return comment, err
	}

	// Query to get the id of the new comment
	row, err := tx.Query("select commentId from comment where photoId = ? and userId = ? and content = ? and time = ?", photoId, comment.UserId, comment.Content, comment.Time)
	if err != nil {
		return comment, err
	}
	if row.Next() {
		err = row.Scan(&comment.CommentId)
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
