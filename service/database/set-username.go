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

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// Query to modify username
	res, err := tx.Exec(`UPDATE user SET username = ? WHERE userId= ?`, newUsername, id)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: user.username" {
			return ErrUsernameAlreadyInUse
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

	// // Query to get the modified user
	// row, err := tx.Query("SELECT * FROM user WHERE ? == userId", id)
	// if err != nil {
	// 	return user, err
	// }
	//
	// if row.Next() {
	// 	err = row.Scan(&user.UserId, &user.Username)
	// 	if err != nil {
	// 		return user, err
	// 	}
	// }
	// defer func() { _ = row.Close() }()

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
