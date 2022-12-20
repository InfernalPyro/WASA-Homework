package database

// Login the user by their username
func (db *appdbimpl) GetStream(id uint64) ([]Photo, error) {
	// All the photos that compone the stream
	var photos []Photo

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return photos, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// Query to get photos of every user that you follow and didn't ban you ordered by the most recent
	row, err := tx.Query("SELECT * FROM photo where userId = (Select follows from follow where userId = ? and follows not in (SELECT userId from ban where banned = ? )) order by time desc", id, id)
	if err != nil {
		return photos, err
	}
	for row.Next() {
		var p Photo
		err = row.Scan(&p.PhotoId, &p.UserId, &p.Image, &p.Time)
		if err != nil {
			return photos, err
		}
		photos = append(photos, p)
	}

	defer func() { _ = row.Close() }()

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return photos, err
	}

	return photos, nil
}
