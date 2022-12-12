package database

// Login the user by their username
func (db *appdbimpl) changeUsername(id int64, newUsername string) (User, error) {
	//Id of the user that is logging in
	var user User

	res, err := db.c.Exec(`UPDATE users SET username=? WHERE id=?`,
		newUsername, id)
	if err != nil {
		return user, err
	}
	row, err := db.c.Query("SELECT * FROM users WHERE ? == userId", id)
	if err != nil {
		return user, err
	}

	for row.Next() {
		err = row.Scan(user)
		if err != nil {
			return user, err
		}
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return user, err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return user, UserDoesNotExist
	}
	return user, nil
}
