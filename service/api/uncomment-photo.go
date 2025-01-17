package api

import (
	"net/http"
	"strconv"

	"github.com/InfernalPyro/WASA-Homework/service/api/reqcontext"
	"github.com/InfernalPyro/WASA-Homework/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// The comment ID that we want to remove in the path is a 64-bit unsigned integer. Let's parse it.
	commentId, err := strconv.ParseUint(ps.ByName("commentId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// The ID of the photo the user want to uncomment in the path is a 64-bit unsigned integer. Let's parse it.
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// We first need to get the owner of the comment we're trying to delete and then we check if the logged user is the owner
	var comment database.Comment
	comment, err = rt.db.GetComment(commentId, photoId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't remove the comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Check if user have permission to make the request
	b, err := Authorized(r.Header.Get("Authorization"), comment.UserId)
	if b == false {
		ctx.Logger.WithError(err).Error("Token error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the function to delete the comment
	err = rt.db.UncommentPhoto(commentId, photoId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't remove the comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.WriteHeader(http.StatusNoContent)
}
