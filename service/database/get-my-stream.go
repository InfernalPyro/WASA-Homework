package database

// Login the user by their username
func (db *appdbimpl) GetStream(id uint64) ([]Photo, string, error) {
	// All the photos that compone the stream
	var photos []Photo
	var username = ""

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return photos, username, err
	}

	// Query to get photos of every user that you follow and didn't ban you ordered by the most recent
	row, err := tx.Query("SELECT photo.*, user.username FROM photo, user where photo.userId = user.userId and photo.userId in (Select follows from follow where userId = ? and follows not in (SELECT userId from ban where banned = ? )) order by time desc", id, id)
	if err != nil {
		return photos, username, err
	}
	for row.Next() {
		var p Photo
		err = row.Scan(&p.PhotoId, &p.UserId, &p.Image, &p.Time, &username)
		if err != nil {
			return photos, username, err
		}
		photos = append(photos, p)
	}
	if err = row.Err(); err != nil {
		return photos, username, err
	}

	defer func() { _ = row.Close() }()

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return photos, username, err
	}

	return photos, username, nil
}
