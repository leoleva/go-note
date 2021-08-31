package util

import (
	"demoproject/src/entity"
	"net/http"
)

func GetUserFromRequest(r *http.Request) entity.User {
	return r.Context().Value("user").(entity.User)
}
