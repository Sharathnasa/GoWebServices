package controller

import (
	"encoding/json"
	"github.com/pluralsight/webservice/models"
	"net/http"
	"regexp"
	"strconv"
)

type userController struct {
	userIdPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/user" {
		switch r.Method {
		case http.MethodGet:
			uc.GetAll(w, r)
		case http.MethodPost:
			uc.post(w, r)

		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIdPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			uc.GetById(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotFound)
		}

	}
}

//get all method
func (uc userController) GetAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJson(models.GetUsers(), w)
}

//get by ID
func (uc *userController) GetById(id int, w http.ResponseWriter) {
	u, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//return in the form of JSON
	encodeResponseAsJson(u, w)
}

//to add the new users to user collection
func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse the object"))
		return
	}

	u, err = models.AddUser(u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJson(u, w)
}

//update the particular user to user collection
func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse the object"))
		return
	}

	if id != u.ID {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ID of the submitted user must match ID in URL"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete the user by userId
func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

//Generic function to parse
func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	//get the JSON body
	dec := json.NewDecoder(r.Body)

	// convert to user model
	var u models.User
	err := dec.Decode(&u)

	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

//new way of creating the constructor(append with new keyword)
// This constructor return the route endpoint
func newUserController() *userController {
	return &userController{userIdPattern: regexp.MustCompile(`^/user/(\d+)/?`)}
}
