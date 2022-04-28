package common

import "net/http"

type Error struct {
	Status int
	Text   string
}

func NewBadRequestError(text string) *Error {
	return &Error{
		Status: http.StatusBadRequest,
		Text:   text,
	}
}

func NewNotImplementedError(text string) *Error {
	return &Error{
		Status: http.StatusNotImplemented,
		Text:   text,
	}
}

func NewNotFoundError(text string) *Error {
	return &Error{
		Status: http.StatusNotFound,
		Text:   text,
	}
}

func (m Error) Error() string {
	return m.Text
}
