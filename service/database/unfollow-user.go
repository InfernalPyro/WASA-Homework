package database

import "errors"

// Make user "id" follow user "followId"
func (db *appdbimpl) UnfollowUser(id uint64, followId uint64) error {

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}

	// Query to remove the follow
	row, err := tx.Exec("DELETE from follow where userId = ? and follows = ?", id, followId)
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
	// If we didn't update any row, then the user didn't exist
	if affected == 0 {
		if tx.Rollback() != nil {
			return err
		}
		return errors.New("Follow does not exists")
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
