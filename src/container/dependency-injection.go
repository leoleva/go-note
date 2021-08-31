package container

import (
	"database/sql"
	"demoproject/src/config"
	"demoproject/src/controller"
	"demoproject/src/middleware"
	"demoproject/src/repository"
	"demoproject/src/util"
	"fmt"
)

type Container struct {
	LoginController *controller.LoginController
	RegisterController *controller.RegisterController
	CreateNoteController *controller.CreateNoteController
	DeleteNoteController *controller.DeleteNoteController
	GetNoteController *controller.GetNoteController

	AuthenticationMiddleware *middleware.AuthenticationMiddleware
}

func Load() *Container {
	cfg, e := loadConfig()
	handleErrors(e)

	database, e := loadDatabase(cfg.Database)
	handleErrors(e)

	userRepository := loadUserRepository(database)
	noteRepository := loadNoteRepository(database)
	loginController := loadLoginController(*userRepository, cfg.JWT)
	registerController := loadRegisterController(*userRepository)
	authenticationMiddleware := loadAuthenticationMiddleware(*userRepository, cfg.JWT)

	return &Container{
		LoginController: loginController,
		RegisterController: registerController,
		AuthenticationMiddleware: authenticationMiddleware,
		CreateNoteController: loadCreateNoteController(*noteRepository),
		DeleteNoteController: loadDeleteNoteController(*noteRepository),
		GetNoteController: loadGetNoteController(*noteRepository),
	}
}

func loadConfig() (*config.Config, error) {
	return config.GetConfig()
}

func loadDatabase(database config.Database) (*sql.DB, error)  {
	return util.GetDatabase(database)
}

func loadUserRepository(db *sql.DB) *repository.UserRepository {
	return repository.NewUserRepo(db)
}

func loadNoteRepository(db *sql.DB) *repository.NoteRepository {
	return repository.NewNoteRepository(db)
}

func loadLoginController(userRepository repository.UserRepository, jwt config.JWT) *controller.LoginController {
	return controller.NewLoginController(userRepository, jwt)
}

func loadRegisterController(userRepository repository.UserRepository) *controller.RegisterController {
	return controller.NewRegisterController(userRepository)
}

func loadAuthenticationMiddleware(userRepository repository.UserRepository, jwt config.JWT) *middleware.AuthenticationMiddleware {
	return middleware.NewAuthenticationMiddleware(userRepository, jwt)
}

func loadCreateNoteController(noteRepository repository.NoteRepository) *controller.CreateNoteController {
	return controller.NewCreateNoteController(noteRepository)
}

func loadDeleteNoteController(noteRepository repository.NoteRepository) *controller.DeleteNoteController {
	return controller.NewDeleteNoteController(noteRepository)
}

func loadGetNoteController(noteRepository repository.NoteRepository) *controller.GetNoteController {
	return controller.NewGetNoteController(noteRepository)
}

// todo: logger to file would be nice
func handleErrors(e error) {
	if e == nil {
		return
	}

	fmt.Println(e)
	panic(e)
}
