// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
)

func (s *ErrorRespStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/Error
type Error struct {
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *Error) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *Error) SetError(val string) {
	s.Error = val
}

// ErrorRespStatusCode wraps Error with StatusCode.
type ErrorRespStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorRespStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorRespStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorRespStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorRespStatusCode) SetResponse(val Error) {
	s.Response = val
}

// Ref: #/components/schemas/Healthy
type Healthy struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *Healthy) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *Healthy) SetMessage(val string) {
	s.Message = val
}

// Ref: #/components/schemas/NewTask
type NewTask struct {
	Name string `json:"name"`
}

// GetName returns the value of Name.
func (s *NewTask) GetName() string {
	return s.Name
}

// SetName sets the value of Name.
func (s *NewTask) SetName(val string) {
	s.Name = val
}

// Ref: #/components/schemas/Task
type Task struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *Task) GetID() int64 {
	return s.ID
}

// GetName returns the value of Name.
func (s *Task) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *Task) SetID(val int64) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Task) SetName(val string) {
	s.Name = val
}

// TasksIDDeleteOK is response for TasksIDDelete operation.
type TasksIDDeleteOK struct{}
