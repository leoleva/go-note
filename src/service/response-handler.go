package service

import (
	"demoproject/src/config"
	"demoproject/src/controller"
	"demoproject/src/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseHandler(c controller.Controller) func(w http.ResponseWriter, r *http.Request) {
	resp := util.NewResponse()

	return func(w http.ResponseWriter, r *http.Request) {
		response, err := c.Handle(r, resp)

		if err != nil {
			fmt.Println("Error from controller: ")
			fmt.Println(err)

			WriteError(w, "Internal system error", 500, config.InternalSystemError)

			return
		}

		if response.GetApiErrorCode() != 0 {
			WriteError(w, response.GetBody(), response.GetStatusCode(), response.GetApiErrorCode())

			return
		}

		write(w, response.GetBody(), response.GetStatusCode())
	}
}

func WriteError(w http.ResponseWriter, output interface{}, statusCode int, apiErrorCode int)  {
	e := make(map[string]interface{})

	e["error"] = output
	e["error_code"] = apiErrorCode

	write(w, e, statusCode)
}

func write(w http.ResponseWriter, output interface{}, statusCode int)  {
	 var toConvert interface{}

	if mappedObject, ok := output.(interface{ToMap() map[string]interface{}}); ok {

		toConvert = mappedObject.ToMap()
	} else if output == nil { // todo: not sure if it's the right way
		toConvert = make(map[string]string)
	} else {
		toConvert = output
	}

	marshal, err := json.Marshal(toConvert)

	if err != nil {
		fmt.Println("Output could not be converted to json")
		fmt.Println(err)

		marshal = []byte("Internal system error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(statusCode)
	_, err = w.Write(marshal)

	if err != nil {
		fmt.Println("Could not write to output")
		fmt.Println(err)

		return
	}
}
