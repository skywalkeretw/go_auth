package controllers

import (
	"github.com/skywalkeretw/auth/middlewares"
	"net/http"
)

func (s *Server) initializeRoutes() {
	// Account
	s.Router.HandleFunc("/register", middlewares.SetMiddlewareJSON(s.Register)).Methods(http.MethodPost)
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods(http.MethodPost)
	// Users
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAdminAuthentication(s.GetUsers))).Methods(http.MethodGet)
	s.Router.HandleFunc("/users/{id:[0-9]+}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAdminAuthentication(s.GetUser))).Methods(http.MethodGet)
	s.Router.HandleFunc("/updateUser", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods(http.MethodPost)
	s.Router.HandleFunc("/deleteUser", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteUser))).Methods(http.MethodPost)
	// Email
	// s.Router.HandleFunc("/confirmUser", middlewares.SetMiddlewareJSON(s.ConfirmUser)).Methods(http.MethodGet)
	// s.Router.HandleFunc("/resetPassword", middlewares.SetMiddlewareJSON(s.ResetPassword)).Methods(http.MethodPost)
}
