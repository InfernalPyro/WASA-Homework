package database

// Make user "id" ban user "banId"
func (db *appdbimpl) BanUser(id uint64, banId uint64) error {

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return err
	}

	// Query to insert the new ban
	_, err = tx.Exec("INSERT into ban (userId,banned) values (?,?)", id, banId)
	if err != nil {
		if err.Error() == ErrForeignKey {
			err = ErrUserNotFound
		}
		if tx.Rollback() != nil {
			return err
		}
		return err
	}

	// If you are following the user you want to ban, remove the follow
	_, err = tx.Exec("DELETE from follow where userId = ? and follows = ?", id, banId)
	if err != nil {
		if tx.Rollback() != nil {
			return err
		}
		return err
	}

	// If the user you want to ban is following you, remove their follow
	_, err = tx.Exec("DELETE from follow where follows = ? and userId = ?", id, banId)
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
