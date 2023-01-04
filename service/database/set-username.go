package database

// Login the user by their username
func (db *appdbimpl) ChangeName(id uint64, newUsername string) error {
	// The user that gets modified
	// var user User

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}

	// Query to modify username
	res, err := tx.Exec(`UPDATE user SET username = ? WHERE userId= ?`, newUsername, id)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: user.username" {
			err = ErrUsernameAlreadyInUse
		}
		if tx.Rollback() != nil {
			return err
		}
		return err
	}

	// Check if the query worked
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	// If we didn't update any row, then the user didn't exist
	if affected == 0 {
		return ErrUserNotFound
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
