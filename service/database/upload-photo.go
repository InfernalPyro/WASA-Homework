package database

import "time"

// Login the user by their username
func (db *appdbimpl) UploadPhoto(id uint64, bImage string) (Photo, error) {

	var photo Photo

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return photo, err
	}

	// Get current time
	layout := "2006-01-02 15:04:05"
	t := time.Now()

	// We will return this photo in database form
	photo.UserId = id
	photo.Image = bImage
	photo.Time = t.Format(layout)

	// Query to modify username
	_, err = tx.Exec(`Insert into photo (userId, image, time) values (?, ?, ?)`, photo.UserId, photo.Image, photo.Time)
	if err != nil {
		if err.Error() == "FOREIGN KEY constraint failed" {
			err = ErrUserNotFound
		}
		if tx.Rollback() != nil {
			return photo, err
		}
		return photo, err
	}

	// Query to get the id of the new photo published
	row, err := tx.Query("select photoId from photo where userId = ? and image = ? and time = ?", photo.UserId, photo.Image, photo.Time)
	if err != nil {
		return photo, err
	}
	if row.Next() {
		err = row.Scan(&photo.PhotoId)
		if err != nil {
			return photo, err
		}
	}
	if err = row.Err(); err != nil {
		return photo, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return photo, err
	}

	return photo, nil
}
