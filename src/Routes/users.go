// Package userRoutes provides ...
package routes

import userHandler "github.com/Raylynd6299/secrets/src/Handlers/User"

func init() {
	UserPath.Path("/info").HandlerFunc(userHandler.UserInfo).Methods("GET")
}
