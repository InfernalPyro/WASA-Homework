package database

// Get photo likes and comments from it's id
func (db *appdbimpl) GetPhotoInfo(photoId uint64) ([]Comment, []Like, error) {

	var comments []Comment
	var likes []Like

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return comments, likes, err
	}

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
	// Query to get the comments
	row, err := tx.Query("SELECT * FROM comment WHERE ? == photoId", photoId)
	if err != nil {
		return comments, likes, err
	}

	for row.Next() {
		var c Comment
		err = row.Scan(&c.CommentId, &c.PhotoId, &c.UserId, &c.Content, &c.Time)
		if err != nil {
			return comments, likes, err
		}
		comments = append(comments, c)
	}
	if err = row.Err(); err != nil {
		return comments, likes, err
	}
	defer func() { _ = row.Close() }()

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
	// Query to get the likes
	row, err = tx.Query("SELECT * FROM like where photoId = ?; ", photoId)
	if err != nil {
		return comments, likes, err
	}
	for row.Next() {
		var l Like
		err = row.Scan(&l.UserId, &l.PhotoId)
		if err != nil {
			return comments, likes, err
		}
		likes = append(likes, l)
	}
	if err = row.Err(); err != nil {
		return comments, likes, err
	}
	defer func() { _ = row.Close() }()

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return comments, likes, err
	}

	return comments, likes, nil
}
