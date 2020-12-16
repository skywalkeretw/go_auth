package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/skywalkeretw/auth/auth"
	"github.com/skywalkeretw/auth/models"
	"github.com/skywalkeretw/auth/responses"
	"io/ioutil"
	"log"
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
	user, token, err := s.SignIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.UserResponse{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		JWT:       token,
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

	err = user.Validate("register")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	log.Println("3: user:",user)
	password := user.Password

	// password is hashed after this point
	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	SendConfirmationEmail(userCreated)

	user, token, err := server.SignIn(user.Email, password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	data := models.UserResponse{
		Firstname:	userCreated.Firstname,
		Lastname:  userCreated.Lastname,
		Email:     userCreated.Email,
		JWT:       token,
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, data)
}

// ConfirmUser new User
func (server *Server) ConfirmUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	claims, err := auth.GetClaims(token)
	if err != nil {
		log.Println(err)
	}
	log.Println("request: ", r)
	log.Println("Claims: ", claims)

	if claims["confirm_user"] != true {
		responses.ERROR(w, http.StatusForbidden, errors.New("Can´t confirm User"))
	}

	//userID, err := strconv.ParseUint(claims["user_id"].(string), 10 , 32)
	userID := uint32(claims["user_id"].(float64))
	log.Printf("%T", userID)

	user := models.User{}
	updatedUser, err := user.ConfirmUser(server.DB, userID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	log.Println("Confirmed User:", updatedUser)

	w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
	responses.JSON(w, http.StatusCreated, models.UserData{
		Fistname: 	updatedUser.Firstname,
		Lastname:  updatedUser.Lastname,
		Email:     updatedUser.Email,
		Type:      updatedUser.Type,
		Confirmed: updatedUser.Confirmed,
	})
}


/*
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