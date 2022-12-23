package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/InfernalPyro/WASA-Homework/service/api/reqcontext"
	"github.com/InfernalPyro/WASA-Homework/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// The User ID that want to like the photo in the path is a 64-bit unsigned integer. Let's parse it.
	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// The ID of the photo the user want to like in the path is a 64-bit unsigned integer. Let's parse it.
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the function make user "id" follow the user "followId"
	err = rt.db.UnlikePhoto(id, photoId)
	if err != nil {
		if errors.Is(err, database.ErrPhotoNotFound) {
			ctx.Logger.WithError(err).Error("Photo does not exists")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		ctx.Logger.WithError(err).Error("Can't remove the like")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.WriteHeader(http.StatusNoContent)
}
