package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/kensei18/enechain-technical-assignment/app/entity"
)

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	IsPrivate   bool       `json:"isPrivate"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`

	CompanyID string
	AuthorID  string
}

func NewTask(task entity.Task) *Task {
	return &Task{
		ID:          task.ID.String(),
		Title:       task.Title,
		Description: task.Description,
		Status:      TaskStatus(task.Status),
		IsPrivate:   task.IsPrivate,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		CompanyID:   task.CompanyID.String(),
		AuthorID:    task.AuthorID.String(),
	}
}

type TaskInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	IsPrivate   bool       `json:"isPrivate"`
	AssigneeIds []string   `json:"assigneeIds"`
}

type TaskUpdateInput struct {
	ID          string      `json:"id"`
	Title       *string     `json:"title,omitempty"`
	Description *string     `json:"description,omitempty"`
	Status      *TaskStatus `json:"status,omitempty"`
	IsPrivate   *bool       `json:"isPrivate,omitempty"`
	AssigneeIds []string    `json:"assigneeIds,omitempty"`
}

type CreateTaskPayload struct {
	Task *Task `json:"task"`
}

type UpdateTaskPayload struct {
	Task *Task `json:"task"`
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
