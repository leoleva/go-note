package router

import (
	"demoproject/src/container"
	"demoproject/src/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RegisterRoutes(container *container.Container)  {
	router := mux.NewRouter().StrictSlash(true)

	setRegisterRoute(container, router)
	setLoginRoute(container, router)
	setCreateNoteRoute(container, router)
	setDeleteNoteRoute(container, router)
	setGetNoteRoute(container, router)

	log.Fatal(http.ListenAndServe(":9090", router))
}

func setRegisterRoute(container *container.Container, router *mux.Router) {
	controller := http.HandlerFunc(service.ResponseHandler(container.RegisterController))

	router.Handle("/register", controller)
}

func setLoginRoute(container *container.Container, router *mux.Router) {
	controller := http.HandlerFunc(service.ResponseHandler(container.LoginController))

	router.Handle("/login", controller)
}

func setCreateNoteRoute(container *container.Container, router *mux.Router) {
	controller := http.HandlerFunc(service.ResponseHandler(container.CreateNoteController))

	router.Handle("/note/create", container.AuthenticationMiddleware.Handler(controller))
}

func setDeleteNoteRoute(container *container.Container, router *mux.Router) {
	controller := http.HandlerFunc(service.ResponseHandler(container.DeleteNoteController))

	router.Handle("/note/delete/{uuid}", container.AuthenticationMiddleware.Handler(controller))
}

func setGetNoteRoute(container *container.Container, router *mux.Router) {
	controller := http.HandlerFunc(service.ResponseHandler(container.GetNoteController))

	router.Handle("/note/get/{uuid}", container.AuthenticationMiddleware.Handler(controller))
}
