package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.NotFound = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	mux.HandlerFunc("GET", "/status", app.status)
	mux.HandlerFunc("POST", "/users", app.createUser)
	mux.HandlerFunc("POST", "/authentication-tokens", app.createAuthenticationToken)

	mux.Handler("GET", "/protected", app.requireAuthenticatedUser(http.HandlerFunc(app.protected)))

	mux.Handler("GET", "/basic-auth-protected", app.requireBasicAuthentication(http.HandlerFunc(app.protected)))

	return app.logAccess(app.recoverPanic(app.authenticate(mux)))
}
