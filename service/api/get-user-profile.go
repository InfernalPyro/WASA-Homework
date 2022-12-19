package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/InfernalPyro/WASA-Homework/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// The User ID in the path is a 64-bit unsigned integer. Let's parse it.
	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the function to change the name
	dbUser, dbFollow, dbFollowed, dbBans, dbComments, dbPhotos, dbLikes, err := rt.db.GetProfile(id)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("Can't get profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// At this point i have a collection of database photos so i need to get, for each photo,
	// it's likes and comments and then convert it into api form
	var photos []Photo
	for _, p := range *dbPhotos {
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

	// Now that every photo of the user is in api form i can build the user in api
	var user User
	user.UserFromDatabase(*dbUser, *dbFollow, *dbFollowed, *dbBans, *dbComments, photos, *dbLikes)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
