package controller

import (
	"database/sql"
	"demoproject/src/config"
	"demoproject/src/repository"
	"demoproject/src/util"
	"github.com/gorilla/mux"
	"net/http"
)

type GetNoteController struct {
	noteRepository repository.NoteRepository
}

func NewGetNoteController(noteRepository repository.NoteRepository) *GetNoteController {
	return &GetNoteController{
		noteRepository: noteRepository,
	}
}

func (c GetNoteController) Handle(r *http.Request, response *util.Response) (*util.Response, error) {
	uuid := mux.Vars(r)["uuid"]
	user := util.GetUserFromRequest(r)

	note, err := c.noteRepository.GetNoteByUserIdAndUuid(user.Id, uuid)

	if err != nil {
		if err == sql.ErrNoRows {
			return response.WithApiError("note not found", 404, config.ResourceNotFound), nil
		}

		return response, err
	}

	return response.WithJson(&note, 200), nil
}
