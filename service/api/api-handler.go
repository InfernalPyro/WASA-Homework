package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	// Login
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	// Change Username
	rt.router.PATCH("/user/:userId/username", rt.wrap(rt.setMyUserName))
	// Get user stream
	rt.router.GET("/user/:userId/home", rt.wrap(rt.getMyStream))
	// Get user profile
	rt.router.GET("/user/:userId/profile", rt.wrap(rt.getUserProfile))
	// Follow user
	rt.router.PUT("/user/:userId/follow/:followId", rt.wrap(rt.followUser))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
