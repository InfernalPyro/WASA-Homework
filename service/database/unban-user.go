package database

import "errors"

// Make user "id" ban user "banId"
func (db *appdbimpl) UnbanUser(id uint64, banId uint64) error {

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}

	// Query to remove the ban
	row, err := tx.Exec("DELETE from ban where userId = ? and banned = ?", id, banId)
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
	// If we didn't update any row, then the ban didn't exist
	if affected == 0 {
		if tx.Rollback() != nil {
			return err
		}
		return errors.New("Ban does not exists")
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
