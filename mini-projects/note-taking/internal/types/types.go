package types

// for creation of notes
type CreateTaskRequest struct {
	NoteTitle       string `json:"note_title" validate:"required,min=5,max=100"`
	NoteDescription string `json:"note_description" validate:"required,min=10,max=100"`
}

type CreateTaskResponse struct {
	Message string
	NoteId  string
}

// for updation of notes
type UpdateTaskRequest struct {
	NoteTitle       string `json:"note_title" validate:"required"`
	NoteDescription string `json:"note_description" validate:"required"`
}

type UpdateTaskResponse struct {
	Message string `json:"message"`
	NoteId  string `json:"note_id"`
}

type DeleteTaskResponse struct {
	Message string
}
