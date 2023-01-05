package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/InfernalPyro/WASA-Homework/service/api/reqcontext"
	"github.com/InfernalPyro/WASA-Homework/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// The ID of the photo the user want to like in the path is a 64-bit unsigned integer. Let's parse it.
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Read the comment from the request body.
	var apiComment Comment
	err = json.NewDecoder(r.Body).Decode(&apiComment)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if user have permission to make the request
	b, err := Authorized(r.Header.Get("Authorization"), apiComment.UserId)
	if b == false {
		ctx.Logger.WithError(err).Error("Token error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Convert ApiComment to database.Comment
	layout := "2006-01-02 15:04:05"
	t := time.Now()

	var dbComment database.Comment
	dbComment.PhotoId = photoId
	dbComment.UserId = apiComment.UserId
	dbComment.Content = apiComment.Comment
	dbComment.Time = t.Format(layout)

	// Call the function make user "id" follow the user "followId"
	newComment, err := rt.db.CommentPhoto(photoId, dbComment)
	if err != nil {
		if errors.Is(err, database.ErrPhotoNotFound) {
			ctx.Logger.WithError(err).Error("Photo does not exists")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		ctx.Logger.WithError(err).Error("Can't add the comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Convert the database comment we just inserted into a struct.Comment for the api
	var comment Comment
	comment.CommentId = newComment.CommentId
	comment.PhotoID = newComment.PhotoId
	comment.UserId = newComment.UserId
	comment.Comment = newComment.Content
	comment.Time = newComment.Time

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)
}
