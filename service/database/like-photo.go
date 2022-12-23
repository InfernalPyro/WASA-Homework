package database

// Make user "id" follow user "followId"
func (db *appdbimpl) LikePhoto(id uint64, photoId uint64) error {

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}

	// Query to insert the new like
	_, err = tx.Exec("INSERT into like (userId,photoId) values (?,?)", id, photoId)
	if err != nil {
		if err.Error() == "FOREIGN KEY constraint failed" {
			err = ErrPhotoNotFound
		}
		if tx.Rollback() != nil {
			return err
		}
		return err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
