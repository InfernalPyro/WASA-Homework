package database

// Login the user by their username
func (db *appdbimpl) LoginUser(username string) (int64, error) {
	// Id of the user that is logging in
	var id int64

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return id, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// Query for the login
	// First a query to check if username already exists. If it exists, return the id.
	row, err := tx.Query("SELECT userId FROM user WHERE EXISTS (SELECT * FROM user WHERE username == ?)", username)
	if err != nil {
		return id, err
	}

	// Check if the query have found the user
	if !row.Next() {
		defer func() { _ = row.Close() }()
		// If not make another query to add the user.
		_, err := tx.Exec(`INSERT INTO user (username) VALUES (?)`, username)
		if err != nil {
			return id, err
		}
	}

	// Then make a query to get the id of this user
	row, err = tx.Query("SELECT userId FROM user WHERE username == ?", username)
	if err != nil {
		return id, err
	}
	if row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return id, err
		}
	}
	defer func() { _ = row.Close() }()

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return id, err
	}

	return id, nil
}
