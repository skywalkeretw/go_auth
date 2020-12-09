package controllers

import (
	"github.com/skywalkeretw/auth/middlewares"
	"net/http"
)

func (s *Server) initializeRoutes() {
	// Account
	s.Router.HandleFunc("/signUp", middlewares.SetMiddlewareJSON(s.Register)).Methods(http.MethodPost)
	s.Router.HandleFunc("/signIn", middlewares.SetMiddlewareJSON(s.Login)).Methods(http.MethodPost)
	// Users
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(middlewares.SetMiddlewareAdminAuthentication(s.GetUsers)))).Methods(http.MethodGet)
}
