package api

import (
	"encoding/json"
	"errors"
	"github.com/InfernalPyro/WASA-Homework/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Read the username from the request body.
	username := ps.ByName("username")

	if r.Header.Get("Authorization") == "" {
		ctx.Logger.WithError(errors.New("Token not found")).Error("Token error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Login the user in the DB.
	dbUsers, err := rt.db.SearchUserByName(username)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("Can't find any user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var users []SimpleUser
	for _, u := range dbUsers {
		var user SimpleUser
		user.UserId = u.UserId
		user.Username = u.Username
		users = append(users, user)
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(users)
}
