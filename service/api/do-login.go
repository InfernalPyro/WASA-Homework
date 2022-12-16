package api

import (
	"encoding/json"
	"net/http"

	"github.com/InfernalPyro/WASA-Homework/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Read the username from the request body.
	var sess Session
	err := json.NewDecoder(r.Body).Decode(&sess)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//if len(sess.username) < 3 || len(sess.username) > 16 {
	//	// Here we validated the username
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	// Login the user in the DB.
	id, err := rt.db.LoginUser(sess.username)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("Can't login the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(id)
}
