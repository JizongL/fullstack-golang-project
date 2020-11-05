package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice" // New import
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders) 
	dynamicMiddleware := alice.New(app.session.Enable,noSurf,app.authenticate)
	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	// Add the requireAuthentication middleware to the chain.
	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippetForm)) // Add the requireAuthentication middleware to the chain.
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippet)) 
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm)) 
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser)) 
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm)) 
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	// Add the requireAuthentication middleware to the chain.
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser)) 
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer)) 
	return standardMiddleware.Then(mux)
	}
