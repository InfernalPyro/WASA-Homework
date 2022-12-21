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

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The User ID in the path is a 64-bit unsigned integer. Let's parse it.
	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Read the new username from the request body.
	var sess Session
	errr := json.NewDecoder(r.Body).Decode(&sess)
	if errr != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len([]rune(sess.Username)) < 3 || len([]rune(sess.Username)) > 16 {
		// Here we validated the username
		ctx.Logger.WithError(err).Error("New Username not valid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the function to change the name
	dbUser, err := rt.db.ChangeName(id, sess.Username)
	if err != nil {
		if errors.Is(err, database.ErrUsernameAlreadyInUse) {
			ctx.Logger.WithError(err).Error("New Username already in use")
			w.WriteHeader(http.StatusNotAcceptable)
		} else if errors.Is(err, database.ErrUserNotFound) {
			ctx.Logger.WithError(err).Error("User not found")
			w.WriteHeader(http.StatusNotFound)
		} else {
			ctx.Logger.WithError(err).Error("Can't get profile")
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// var user User
	// user.FromDatabase(dbUser)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(dbUser)
}
