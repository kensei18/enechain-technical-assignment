package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	IsPrivate   bool       `json:"isPrivate"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type TaskInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	IsPrivate   bool       `json:"isPrivate"`
}

type TaskUpdateInput struct {
	ID          string      `json:"id"`
	Title       *string     `json:"title,omitempty"`
	Description *string     `json:"description,omitempty"`
	Status      *TaskStatus `json:"status,omitempty"`
	IsPrivate   *bool       `json:"isPrivate,omitempty"`
}

type CreateTaskPayload struct {
	Tasks []*Task `json:"tasks"`
}

type UpdateTaskPayload struct {
	Tasks []*Task `json:"tasks"`
}

type TaskStatus string

const (
	TaskStatusTodo    TaskStatus = "Todo"
	TaskStatusOnGoing TaskStatus = "OnGoing"
	TaskStatusDone    TaskStatus = "Done"
)

var AllTaskStatus = []TaskStatus{
	TaskStatusTodo,
	TaskStatusOnGoing,
	TaskStatusDone,
}

func (e TaskStatus) IsValid() bool {
	switch e {
	case TaskStatusTodo, TaskStatusOnGoing, TaskStatusDone:
		return true
	}
	return false
}

func (e TaskStatus) String() string {
	return string(e)
}

func (e *TaskStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskStatus", str)
	}
	return nil
}

func (e TaskStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
