package database

// With this function we can get the id and username of every user that contains the username
func (db *appdbimpl) SearchUserByName(username string) ([]User, error) {
	// List of ids
	var users []User
	username = "%" + username + "%"

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return users, err
	}

	// Make a query to get the id of this user
	row, err := tx.Query("SELECT * FROM user WHERE username LIKE ?", username)
	if err != nil {
		return users, err
	}
	defer func() { _ = row.Close() }()

	for row.Next() {
		var u User
		err = row.Scan(&u.UserId, &u.Username)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	if err = row.Err(); err != nil {
		return users, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return users, err
	}

	return users, nil
}
