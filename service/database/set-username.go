package database

import "log"

// Login the user by their username
func (db *appdbimpl) ChangeName(id uint64, newUsername string) (User, error) {
	//The user that gets modified
	var user User

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return user, err
	}

	//Query to modify username
	res, err := tx.Exec(`UPDATE user SET username = ? WHERE userId= ?`, newUsername, id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update drivers: unable to rollback: %v", rollbackErr)
		}
		return user, err
	}

	//Check if the query worked
	affected, err := res.RowsAffected()
	if err != nil {
		return user, err
	}
	// If we didn't update any row, then the user didn't exist
	if affected == 0 {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update drivers: unable to rollback: %v", rollbackErr)
		}
		return user, UserDoesNotExist
	}

	//Query to get the modified user
	row, err := tx.Query("SELECT * FROM user WHERE ? == userId", id)
	if err != nil {
		return user, err
	}

	if row.Next() {
		err = row.Scan(&user.UserId, &user.Username)
		if err != nil {
			return user, err
		}
	}
	defer func() { _ = row.Close() }()

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return user, err
	}

	return user, nil
}
