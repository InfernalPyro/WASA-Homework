package database

// Make user "id" follow user "followId"
func (db *appdbimpl) FollowUser(id uint64, followId uint64) error {

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}

	// First a query to check if the user is banned from the one that they want to follow
	row, err := tx.Query("SELECT * FROM ban WHERE banned == ? and userId == ?", id, followId)
	if err != nil {
		return err
	}
	defer func() { _ = row.Close() }()

	// Check if the query have found the user
	if row.Next() {
		if tx.Rollback() != nil {
			return err
		}
		return ErrUserIsBanned
	}
	if err = row.Err(); err != nil {
		return err
	}

	// Query to insert the new follow
	_, err = tx.Exec("INSERT into follow (userId,follows) values (?,?)", id, followId)
	if err != nil {
		if err.Error() == "FOREIGN KEY constraint failed" {
			err = ErrUserNotFound
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
