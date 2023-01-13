package api

import (
	"errors"
	"github.com/InfernalPyro/WASA-Homework/service/database"
	"strconv"
	"strings"
)

type Session struct {
	Username string `json:"username"`
}

type Image struct {
	Image string `json:"image"`
}

type SimpleUser struct {
	UserId    uint64 `json:"id"`
	Username string `json:"username"`
}

type Comment struct {
	CommentId uint64 `json:"commentId"`
	PhotoID   uint64 `json:"photoId"`
	UserId    uint64 `json:"userId"`
	Comment   string `json:"comment"`
	Time      string `json:"time"`
}

type Photo struct {
	PhotoId   uint64    `json:"photoId"`
	ProfileId uint64    `json:"profileId"`
	Image     string    `json:"image"`
	Time      string    `json:"time"`
	Likes     []uint64  `json:"likes"`
	Comments  []Comment `json:"comments"`
}

type User struct {
	UserId    uint64    `json:"userId"`
	Username  string    `json:"username"`
	Follows   []uint64  `json:"follows"`
	Followers []uint64  `json:"followers"`
	Banned    []uint64  `json:"banned"`
	Photos    []Photo   `json:"photos"`
	Likes     []uint64  `json:"likes"`
	Comments  []Comment `json:"comments"`
}

// This function convert all the data of a single user taken from the db into a single user in api form
func (u *User) UserFromDatabase(user database.User, follow []database.Follow, followed []database.Follow, bans []database.Ban, comments []database.Comment, photos []Photo, likes []database.Like) {
	u.UserId = user.UserId
	u.Username = user.Username
	for _, f := range follow {
		u.Follows = append(u.Follows, f.Follows)
	}
	for _, fd := range followed {
		u.Followers = append(u.Followers, fd.UserId)
	}
	for _, b := range bans {
		u.Banned = append(u.Banned, b.Banned)
	}
	for _, p := range photos {
		u.Photos = append(u.Photos, p)
	}
	for _, l := range likes {
		u.Likes = append(u.Likes, l.PhotoId)
	}
	for _, c := range comments {
		var comment Comment
		comment.CommentId = c.CommentId
		comment.PhotoID = c.PhotoId
		comment.UserId = c.UserId
		comment.Comment = c.Content
		comment.Time = c.Time
		u.Comments = append(u.Comments, comment)
	}
	return
}

// This function convert all the data of a single photo taken from the db into a single photo in api form
func (p *Photo) PhotoFromDatabase(photo database.Photo, comments []database.Comment, likes []database.Like) {
	p.PhotoId = photo.PhotoId
	p.ProfileId = photo.UserId
	p.Image = string(photo.Image[:])
	p.Time = photo.Time
	for _, l := range likes {
		p.Likes = append(p.Likes, l.UserId)
	}
	for _, c := range comments {
		var comment Comment
		comment.CommentId = c.CommentId
		comment.PhotoID = c.PhotoId
		comment.UserId = c.UserId
		comment.Comment = c.Content
		comment.Time = c.Time
		p.Comments = append(p.Comments, comment)
	}
	return
}

// This function convert an api photo to a database one
func PhotoToDatabase(photo Photo) database.Photo {
	var p database.Photo
	p.PhotoId = photo.PhotoId
	p.UserId = photo.ProfileId
	p.Image = []byte(photo.Image)
	p.Time = photo.Time

	return p
}

// Function to check if id is authorized to make the operation
func Authorized(bearer string, id uint64) (bool, error) {
	// First get the token
	if bearer == "" {
		return false, errors.New("Token not found")
	}
	// And then check if the logged user id is the same as the token
	token := strings.Split(bearer, " ")
	if token[1] != strconv.FormatUint(id, 10) {
		return false, errors.New("User does not have permission")
	}
	return true, nil
}
