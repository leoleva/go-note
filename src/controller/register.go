package controller

import (
	"demoproject/src/config"
	"demoproject/src/entity"
	"demoproject/src/repository"
	"demoproject/src/util"
	"errors"
	"github.com/google/uuid"
	"net/http"
)

type RegisterController struct {
	userRepo repository.UserRepository
}

func NewRegisterController(userRepo repository.UserRepository) *RegisterController {
	return &RegisterController{
		userRepo: userRepo,
	}
}

func (c *RegisterController) Handle(r *http.Request, response *util.Response) (resp *util.Response, e error) {
	if r.Method != "POST" {
		return response.WithApiError("this endpoint only supports POST calls", 400, config.NotSupportedMethod), nil
	}

	formError := r.ParseForm()

	if formError != nil {
		return response, formError
	}

	email, password, e := getData(r)

	if e != nil {
		return response.WithApiError(e.Error(), 400, config.InvalidFormParameters), nil
	}

	if c.userRepo.UserExists(email) {
		return response.WithApiError("User already exists", 409, config.UserAlreadyExists), nil
	}

	err := createUser(c, email, password)

	if err != nil {
		return nil, err
	}

	return response.WithStatusCode(201), nil
}

func getData(r *http.Request) (email string, password string, err error)  {
	email = r.FormValue("email")
	password = r.FormValue("password")
	passwordRepeat := r.FormValue("password_repeat")

	if email == "" {
		return email, password, errors.New("email is empty")
	}

	if password == "" {
		return email, password, errors.New("password is empty")
	}

	if passwordRepeat == "" {
		return email, password, errors.New("password repeat is empty")
	}

	if passwordRepeat != password {
		return email, password, errors.New("password and password repeat must be equal")
	}

	return email, password, nil
}

func createUser(c *RegisterController, email string, password string) error {
	userUuid, err := uuid.NewRandom()

	if err != nil {
		return nil
	}

	user := entity.User {
		Email: email,
		Password:  password,
		Uuid: userUuid.String(),
		//CreatedAt: time.Now(),
	}

	_, e := c.userRepo.Create(user)

	if e != nil {
		return e
	}

	return nil
}
