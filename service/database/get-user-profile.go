package database

import (
	"errors"
)

// Login the user by their username
func (db *appdbimpl) GetProfile(id uint64) (*User, *[]Follow, *[]Follow, *[]Ban, *[]Comment, *[]Photo, *[]Like, error) {

	var user User
	var follow []Follow
	var followed []Follow
	var bans []Ban
	var comments []Comment
	var photos []Photo
	var likes []Like

	// Get a Tx for making transaction requests.
	tx, err := db.c.Begin()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
	// Query to get the user
	row, err := tx.Query("SELECT * FROM user WHERE ? == userId", id)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	if row.Next() {
		err = row.Scan(&user.UserId, &user.Username)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, nil, err
		}
	} else {
		return nil, nil, nil, nil, nil, nil, nil, errors.New("User does not exists")
	}
	defer func() { _ = row.Close() }()

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
	// Query to get the follows
	row, err = tx.Query("SELECT * FROM follow where userId = ?; ", id)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	for row.Next() {
		var f Follow
		err = row.Scan(&f.UserId, &f.Follows)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, nil, err
		}
		follow = append(follow, f)
	}
	if err = row.Err(); err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}
	defer func() { _ = row.Close() }()

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
	// Query to get the users that follows you
	row, err = tx.Query("SELECT * FROM follow where follows = ?; ", id)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	for row.Next() {
		var fd Follow
		err = row.Scan(&fd.UserId, &fd.Follows)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, nil, err
		}
		followed = append(followed, fd)
	}
	if err = row.Err(); err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}
	defer func() { _ = row.Close() }()

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
	// Query to get the banned
	row, err = tx.Query("SELECT * FROM ban where userId = ?; ", id)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	for row.Next() {
		var b Ban
		err = row.Scan(&b.UserId, &b.Banned)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, nil, err
		}
		bans = append(bans, b)
	}

	if err = row.Err(); err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}
	defer func() { _ = row.Close() }()

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
	// Query to get the comments
	row, err = tx.Query("SELECT * FROM comment where userId = ?; ", id)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	for row.Next() {
		var c Comment
		err = row.Scan(&c.CommentId, &c.PhotoId, &c.UserId, &c.Content, &c.Time)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, nil, err
		}
		comments = append(comments, c)
	}
	if err = row.Err(); err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}
	defer func() { _ = row.Close() }()

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
	// Query to get the photos
	row, err = tx.Query("SELECT photoId,userId,hex(image),time FROM photo where userId = ?; ", id)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	for row.Next() {
		var p Photo
		err = row.Scan(&p.PhotoId, &p.UserId, &p.Image, &p.Time)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, nil, err
		}
		photos = append(photos, p)
	}
	if err = row.Err(); err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}
	defer func() { _ = row.Close() }()

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
	// Query to get the likes
	row, err = tx.Query("SELECT * FROM like where userId = ?; ", id)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}
	for row.Next() {
		var l Like
		err = row.Scan(&l.UserId, &l.PhotoId)
		if err != nil {
			return nil, nil, nil, nil, nil, nil, nil, err
		}
		likes = append(likes, l)
	}
	if err = row.Err(); err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}
	defer func() { _ = row.Close() }()

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	return &user, &follow, &followed, &bans, &comments, &photos, &likes, nil

}
