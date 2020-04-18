package controllers

import (
	"encoding/json"
	"github.com/pluralsight-fundamentals/webservice/models"
	"net/http"
	"regexp"
	"strconv"
)

type UserController struct {
	userIdPattern *regexp.Regexp
}

func (controller UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//_, _ = w.Write([]byte("Hello From UserController"))
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			controller.getALL(w, r)
			break
		case http.MethodPost:
			controller.post(w, r)
			break
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := controller.userIdPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			controller.get(id, w)
			break
		case http.MethodPut:
			controller.put(id, r, w)
			break
		case http.MethodDelete:
			controller.delete(id, w)
			break
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (controller *UserController) getALL(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (controller *UserController) get(id int, w http.ResponseWriter) {
	user, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(user, w)
}

func (controller *UserController) post(w http.ResponseWriter, r *http.Request) {
	user, err := controller.ParseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Could not Parse User Object"))
		return
	}
	user, err = models.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Could not Add User"))
		return
	}
	encodeResponseAsJSON(user, w)
}

func (controller *UserController) put(id int, r *http.Request, w http.ResponseWriter) {
	user, err := controller.ParseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Could not Parse Object "))
		return
	}
	if id != user.ID {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("ID of submitted user must match ID in URL "))
		return
	}
	encodeResponseAsJSON(user, w)
}

func (controller *UserController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (controller *UserController) ParseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var user models.User
	err := dec.Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Constructor
func newUserController() *UserController {
	return &UserController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
