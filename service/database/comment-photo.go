package database

// Make user "id" follow user "followId"
func (db *appdbimpl) CommentPhoto(photoId uint64, comment Comment) (Comment, string, error) {

	var username = ""

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return comment, username, err
	}

	// Query to insert the new like
	_, err = tx.Exec("Insert into comment (photoId, userId, content, time) values (?,?,?,?)", photoId, comment.UserId, comment.Content, comment.Time)
	if err != nil {
		if err.Error() == ErrForeignKey {
			err = ErrPhotoNotFound
		}
		if tx.Rollback() != nil {
			return comment, username, err
		}
		return comment, username, err
	}

	// Query to get the id of the new comment
	row, err := tx.Query("select comment.commentId, user.username from comment,user where photoId = ? and comment.userId = ? and content = ? and time = ? and comment.userId = user.userId", photoId, comment.UserId, comment.Content, comment.Time)
	if err != nil {
		return comment, username, err
	}
	if row.Next() {
		err = row.Scan(&comment.CommentId, &username)
		if err != nil {
			return comment, username, err
		}
	}
	if err = row.Err(); err != nil {
		return comment, username, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return comment, username, err
	}

	return comment, username, nil
}
