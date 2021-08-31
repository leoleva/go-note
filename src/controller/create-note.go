package controller

import (
	"demoproject/src/config"
	"demoproject/src/entity"
	"demoproject/src/repository"
	"demoproject/src/util"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type CreateNoteController struct {
	noteRepository repository.NoteRepository
}

func NewCreateNoteController(noteRepository repository.NoteRepository) *CreateNoteController {
	return &CreateNoteController{
		noteRepository: noteRepository,
	}
}

func (c CreateNoteController) Handle(r *http.Request, response *util.Response) (resp *util.Response, e error) {
	validationErrors := validateCreateNoteRequest(r)

	if validationErrors != nil {
		return resp.WithApiError(validationErrors.Error(), 400, config.InvalidFormParameters), nil
	}

	noteUuid, err := uuid.NewRandom()

	if err != nil {
		return response, err
	}

	note := entity.NewNote(
		r.FormValue("title"),
		r.FormValue("text"),
		getUser(r).Id,
		time.Now(),
		noteUuid.String(),
		)

	createdNote, err := c.noteRepository.Create(*note)

	if err != nil {
		return nil, err
	}

	return response.WithJson(&createdNote,201), nil
}

func validateCreateNoteRequest(r *http.Request) error {
	title := r.FormValue("title")

	if title == "" {
		return errors.New("title input is missing")
	}

	return nil
}

func getUser(r *http.Request) entity.User {
	return r.Context().Value("user").(entity.User)
}
