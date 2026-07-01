package types

type CreateTaskRequest struct {
	NoteTitle       string `json:"note_title" validate:"required,min=5,max=100"`
	NoteDescription string `json:"note_description" validate:"required,min=10,max=100"`
}

type CreateTaskResponse struct {
	Message string
	Data    CreateTaskRequest
}
