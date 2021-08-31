package controller

import (
	"demoproject/src/util"
	"net/http"
)

type Controller interface {
	Handle(r *http.Request, response *util.Response) (resp *util.Response, e error)
}
