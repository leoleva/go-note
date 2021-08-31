package controller

import (
	"database/sql"
	"demoproject/src/config"
	"demoproject/src/repository"
	"demoproject/src/util"
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
)

type LoginController struct {
	repository.UserRepository
	config.JWT
}

func NewLoginController(userRepository repository.UserRepository, config config.JWT) *LoginController {
	return &LoginController {
		UserRepository: userRepository,
		JWT: config,
	}
}

// todo refactor to functions/services

func (c LoginController) Handle(r *http.Request, response *util.Response) (resp *util.Response, e error) {
	email, password := getValues(r)

	validationError := validate(email, password)

	if validationError != nil {
		return response.WithApiError(validationError.Error(), 400, config.InvalidFormParameters), nil
	}

	user, repositoryError := c.UserRepository.GetUserByEmailAndPassword(email, password)

	if repositoryError != nil {
		if repositoryError == sql.ErrNoRows {
			return response.WithApiError("Login failed. Bad email/password combination", 400, config.LoginFailed), nil
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "demoproject",
/*		"exp": time.Now().Add(time.Hour * 24),
*/		"uuid": user.Uuid,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(c.JWT.SecretKey))

	if err != nil {
		return response, err
	}

	jwtResponse := make(map[string]string)

	jwtResponse["jwt"] = tokenString

	return response.WithJson(jwtResponse, 200), nil
}

func validate(email string, password string) error {
	if email == "" {
		return errors.New("email input is missing")
	}

	if password == "" {
		return errors.New("password input is missing")
	}

	return nil
}

func getValues(r *http.Request) (email string, password string)  {
	email = r.FormValue("email")
	password = r.FormValue("password")

	return email, password
}
