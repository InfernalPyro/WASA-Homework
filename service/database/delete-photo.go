package database

import "errors"

// Make user "id" follow user "followId"
func (db *appdbimpl) DeletePhoto(userId uint64, photoId uint64) error {

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}

	// Query to remove the photo
	row, err := tx.Exec("DELETE from photo where userId = ? and photoId = ?", userId, photoId)
	if err != nil {
		if tx.Rollback() != nil {
			return err
		}
		return err
	}

	// Check if the query worked
	affected, err := row.RowsAffected()
	if err != nil {
		if tx.Rollback() != nil {
			return err
		}
		return err
	}
	// If we didn't update any row, then the photo didn't exist
	if affected == 0 {
		if tx.Rollback() != nil {
			return err
		}
		return errors.New("Photo does not exists")
	}

	// Query to remove the likes on the deleted photo
	_, err = tx.Exec("DELETE from like where photoId = ?", photoId)
	if err != nil {
		if tx.Rollback() != nil {
			return err
		}
		return err
	}

	// Query to remove the comments on the deleted photo
	_, err = tx.Exec("DELETE from comment where photoId = ?", photoId)
	if err != nil {
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
