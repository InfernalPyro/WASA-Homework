package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/InfernalPyro/WASA-Homework/service/api/reqcontext"
	"github.com/InfernalPyro/WASA-Homework/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The User ID in the path is a 64-bit unsigned integer. Let's parse it.
	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Read the image from the request body.
	var bImage string
	err = json.NewDecoder(r.Body).Decode(&bImage)
	if err != nil {
		// The body was not a parseable JSON, reject it
		// IL problema è qui, non riesco a leggere l'immagine passata dal request body
		ctx.Logger.WithError(err).Error("User not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Call the function to change the name
	dbPhoto, err := rt.db.UploadPhoto(id, bImage)
	if err != nil {
		if errors.Is(err, database.ErrUserNotFound) {
			ctx.Logger.WithError(err).Error("User not found")
			w.WriteHeader(http.StatusNotFound)
		} else {
			ctx.Logger.WithError(err).Error("Can't publish photo")
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// The photo just got published so it surely won't have comments nor likes
	var comments []database.Comment
	var likes []database.Like
	var photo Photo
	// Convert the photo from database to api form
	photo.PhotoFromDatabase(dbPhoto, comments, likes)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)
}
