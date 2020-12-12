package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/skywalkeretw/auth/models"
	"github.com/skywalkeretw/auth/responses"
	"io/ioutil"
	"net/http"
)

// Login existing User
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := s.SignIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.UserResponse{
		Email: user.Email,
		JWT:   token,
	}
	responses.JSON(w, http.StatusOK, data)
}

// Register new User
func (server *Server) Register(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.UserResponse{
		Email: userCreated.Email,
		JWT:   token,
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, data)
}
/*
// ConfirmUser new User
func (server *Server) ConfirmUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	claims, err := auth.GetClaims(token)
	if err != nil {
		log.Println(err)
	}
	log.Println(claims)

	w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
	responses.JSON(w, http.StatusCreated, "")
}

// ResetPassword sends a confirmation email to the user to reset the password
func (server *Server) ResetPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}


	SendPasswordResetMail(user.Email)
	responses.JSON(w, http.StatusOK, "")
}
*/