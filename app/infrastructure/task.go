package infrastructure

import (
	"context"

	"github.com/google/uuid"
	"github.com/kensei18/enechain-technical-assignment/app/domain/repository"
	"github.com/kensei18/enechain-technical-assignment/app/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type taskRepository struct {
	conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &taskRepository{conn: conn}
}

func (r *taskRepository) Create(ctx context.Context, params repository.CreateTaskParams) (*entity.Task, error) {
	taskAssignees := make([]entity.TaskAssignee, 0, len(params.AssigneeIDs))
	for _, assgineeID := range params.AssigneeIDs {
		taskAssignees = append(taskAssignees, entity.TaskAssignee{AssigneeID: assgineeID})
	}
	task := &entity.Task{
		CompanyID:     params.CompanyID,
		AuthorID:      params.AuthorID,
		Title:         params.Title,
		Description:   params.Description,
		Status:        params.Status,
		IsPrivate:     params.IsPrivate,
		TaskAssignees: taskAssignees,
	}
	if err := r.conn.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *taskRepository) Update(ctx context.Context, params repository.UpdateTaskParams) (*entity.Task, error) {
	values := map[string]interface{}{}
	if params.Title != nil {
		values["title"] = *params.Title
	}
	if params.Description != nil {
		values["description"] = *params.Description
	}
	if params.Status != nil {
		values["status"] = *params.Status
	}
	if params.IsPrivate != nil {
		values["is_private"] = *params.IsPrivate
	}

	task := &entity.Task{ID: params.TaskID}
	if err := r.conn.Model(task).Clauses(clause.Returning{}).Updates(values).Error; err != nil {
		return nil, err
	}

	if params.AssigneeIDs == nil {
		return task, nil
	}
	taskAssignees := make([]entity.TaskAssignee, 0, len(params.AssigneeIDs))
	for _, assigneeID := range params.AssigneeIDs {
		taskAssignees = append(taskAssignees, entity.TaskAssignee{AssigneeID: assigneeID})
	}
	if err := r.conn.Where("task_id = ?", task.ID).Delete(&entity.TaskAssignee{}).Error; err != nil {
		return nil, err
	}
	if err := r.conn.Model(task).Association("TaskAssignees").Append(taskAssignees); err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.conn.Delete(&entity.Task{}, id).Error
}
