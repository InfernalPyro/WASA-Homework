package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/InfernalPyro/WASA-Homework/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// The User ID in the path is a 64-bit unsigned integer. Let's parse it.
	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the function to change the name
	dbPhotos, err := rt.db.GetStream(id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Logger.WithError(err).Error("User not found")
			w.WriteHeader(http.StatusNotFound)
		} else {
			ctx.Logger.WithError(err).Error("Can't get stream")
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// At this point i have a collection of database photos so i need to get, for each photo,
	// it's likes and comments and then convert it into api form
	var photos []Photo
	for _, p := range dbPhotos {
		var photo Photo
		comments, likes, err := rt.db.GetPhotoInfo(p.PhotoId)
		if err != nil {
			ctx.Logger.WithError(err).Error("Can't get photo info")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		photo.PhotoFromDatabase(p, comments, likes)
		photos = append(photos, photo)
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photos)
}
