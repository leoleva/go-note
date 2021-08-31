package controller

import (
	"demoproject/src/config"
	"demoproject/src/repository"
	"demoproject/src/util"
	"github.com/gorilla/mux"
	"net/http"
)

type DeleteNoteController struct {
	noteRepository repository.NoteRepository
}

func NewDeleteNoteController(noteRepository repository.NoteRepository) *DeleteNoteController {
	return &DeleteNoteController{
		noteRepository: noteRepository,
	}
}

func (c DeleteNoteController) Handle(r *http.Request, response *util.Response) (*util.Response, error) {
	uuid := mux.Vars(r)["uuid"]
	user := util.GetUserFromRequest(r)

	if uuid == "" {
		return response.WithApiError("uuid is mandatory", 400, config.InvalidFormParameters), nil
	}

	if c.noteRepository.NoteExistsByUserIdAndUuid(user.Id, uuid) == false {
		return response.WithApiError("note doesnt exist", 400, config.ResourceNotFound), nil
	}

	err := c.noteRepository.DeleteByUserIdAndUuid(user.Id, uuid)

	if err != nil {
		return response, err
	}

	return response.WithStatusCode(200), nil
}
