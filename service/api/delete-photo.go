package api

import (
	"net/http"
	"strconv"

	"github.com/InfernalPyro/WASA-Homework/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// The user ID that published the photo we want to remove in the path is a 64-bit unsigned integer. Let's parse it.
	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// The ID of the photo we want to remove in the path is a 64-bit unsigned integer. Let's parse it.
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if user have permission to make the request
	b, err := Authorized(r.Header.Get("Authorization"), id)
	if b == false {
		ctx.Logger.WithError(err).Error("Token error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the function to delete the comment
	err = rt.db.DeletePhoto(id, photoId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't remove the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.WriteHeader(http.StatusNoContent)
}
