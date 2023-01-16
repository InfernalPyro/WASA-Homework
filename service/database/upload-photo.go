package database

import "time"

// Login the user by their username
func (db *appdbimpl) UploadPhoto(id uint64, photo Photo) (Photo, string, error) {

	var username = ""

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return photo, username, err
	}

	// Get current time
	layout := "2006-01-02 15:04:05"
	t := time.Now()

	// We will return this photo in database form
	photo.UserId = id
	photo.Time = t.Format(layout)

	// Query to modify username
	_, err = tx.Exec(`Insert into photo (userId, image, time) values (?, ?, ?)`, photo.UserId, photo.Image, photo.Time)
	if err != nil {
		if err.Error() == ErrForeignKey {
			err = ErrUserNotFound
		}
		if tx.Rollback() != nil {
			return photo, username, err
		}
		return photo, username, err
	}

	// Query to get the id of the new photo published
	row, err := tx.Query("select photoId,user.username from photo,user where photo.userId = user.userId and photo.userId = ? and image = ? and time = ?", photo.UserId, photo.Image, photo.Time)
	if err != nil {
		return photo, username, err
	}
	if row.Next() {
		err = row.Scan(&photo.PhotoId, &username)
		if err != nil {
			return photo, username, err
		}
	}
	if err = row.Err(); err != nil {
		return photo, username, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return photo, username, err
	}

	return photo, username, nil
}
