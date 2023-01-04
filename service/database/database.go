/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

type Photo struct {
	PhotoId uint64
	UserId  uint64
	Image   string
	Time    string
}

type User struct {
	UserId   uint64
	Username string
}

type Follow struct {
	UserId  uint64
	Follows uint64
}

type Ban struct {
	UserId uint64
	Banned uint64
}

type Comment struct {
	CommentId uint64
	UserId    uint64
	PhotoId   uint64
	Content   string
	Time      string
}

type Like struct {
	UserId  uint64
	PhotoId uint64
}

var ErrUserNotFound = errors.New("User not found")
var ErrPhotoNotFound = errors.New("Photo not found")
var ErrUsernameAlreadyInUse = errors.New("Username alrady in use")
var ErrUserIsBanned = errors.New("User is banned from the other")

const ErrForeignKey = "FOREIGN KEY constraint failed"

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	LoginUser(username string) (int64, error)
	ChangeName(id uint64, newUsername string) error
	GetStream(id uint64) ([]Photo, error)
	GetProfile(id uint64) (*User, *[]Follow, *[]Follow, *[]Ban, *[]Comment, *[]Photo, *[]Like, error)
	GetPhotoInfo(photoId uint64) ([]Comment, []Like, error)
	FollowUser(id uint64, followId uint64) error
	UnfollowUser(id uint64, followId uint64) error
	BanUser(id uint64, banId uint64) error
	UnbanUser(id uint64, banId uint64) error
	LikePhoto(id uint64, photoId uint64) error
	UnlikePhoto(id uint64, photoId uint64) error
	CommentPhoto(photoId uint64, comment Comment) (Comment, error)
	UncommentPhoto(commentId uint64, photoId uint64) error
	GetComment(commentId uint64, photoId uint64) (Comment, error)
	UploadPhoto(id uint64, bImage Photo) (Photo, error)
	DeletePhoto(id uint64, photoId uint64) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	_, err := db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='user';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `create table user(
			userId INTEGER PRIMARY KEY, 
			username TEXT NOT NULL UNIQUE 
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='follow';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `create table follow(
			userId INTEGER, 
			follows INTEGER,
			   PRIMARY KEY (userId, follows),
			   FOREIGN KEY (userId) 
				REFERENCES user (userId) 
					ON DELETE CASCADE 
					ON UPDATE CASCADE,
			   FOREIGN KEY (follows) 
					  REFERENCES user (userId) 
						 ON DELETE CASCADE 
						 ON UPDATE CASCADE
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='ban';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `create table ban(
			userId INTEGER, 
			banned INTEGER,
			   PRIMARY KEY (userId, banned),
			   FOREIGN KEY (userId) 
				REFERENCES user (userId) 
					ON DELETE CASCADE 
					ON UPDATE CASCADE,
			   FOREIGN KEY (banned) 
					  REFERENCES user (userId) 
						 ON DELETE CASCADE 
						 ON UPDATE CASCADE
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='photo';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `create table photo(
			photoId INTEGER PRIMARY KEY, 
			userId INTEGER NOT NULL,
			image BLOB,
			time TEXT,
			FOREIGN KEY (userId) 
				REFERENCES user (userId) 
					ON DELETE CASCADE 
					ON UPDATE CASCADE
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comment';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `create table comment(
			commentId INTEGER PRIMARY KEY, 
			photoId INTEGER NOT NULL, 
			userId INTEGER NOT NULL,
			content TEXT,
			time TEXT,
			FOREIGN KEY (userId) 
				REFERENCES user (userId) 
					ON DELETE CASCADE 
					ON UPDATE CASCADE,
			   FOREIGN KEY (photoId) 
					  REFERENCES photo (photoId) 
						 ON DELETE CASCADE 
						 ON UPDATE CASCADE
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='like';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `create table like(
			userId INTEGER, 
			photoId INTEGER,
			   PRIMARY KEY (userId, photoId),
			   FOREIGN KEY (userId) 
				REFERENCES user (userId) 
					ON DELETE CASCADE 
					ON UPDATE CASCADE,
			   FOREIGN KEY (photoId) 
					  REFERENCES photo (photoId) 
						 ON DELETE CASCADE 
						 ON UPDATE CASCADE
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
