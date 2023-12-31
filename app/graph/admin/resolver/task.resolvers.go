package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/google/uuid"
	"github.com/kensei18/enechain-technical-assignment/app/domain/repository"
	"github.com/kensei18/enechain-technical-assignment/app/domain/service"
	"github.com/kensei18/enechain-technical-assignment/app/entity"
	"github.com/kensei18/enechain-technical-assignment/app/graph/admin"
	"github.com/kensei18/enechain-technical-assignment/app/graph/admin/model"
	"github.com/kensei18/enechain-technical-assignment/app/infrastructure"
	"github.com/kensei18/enechain-technical-assignment/app/storage"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.TaskInput) (*model.CreateTaskPayload, error) {
	authorID, err := uuid.Parse(input.AuthorID)
	if err != nil {
		return nil, err
	}

	status, err := entity.ParseTaskStatus(string(input.Status))
	if err != nil {
		return nil, err
	}

	assigneeIDs := make([]uuid.UUID, 0, len(input.AssigneeIds))
	for _, assigneeID := range input.AssigneeIds {
		id, err := uuid.Parse(assigneeID)
		if err != nil {
			return nil, err
		}
		assigneeIDs = append(assigneeIDs, id)
	}

	taskService := service.NewTaskService(r.DB(ctx))
	task, err := taskService.CreateTask(ctx, service.CreateTaskParams{
		AuthorID:    authorID,
		Title:       input.Title,
		Description: input.Description,
		Status:      status,
		IsPrivate:   input.IsPrivate,
		AssigneeIDs: assigneeIDs,
	})
	if err != nil {
		return nil, err
	}

	return &model.CreateTaskPayload{Task: model.NewTask(*task)}, nil
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, input model.TaskUpdateInput) (*model.UpdateTaskPayload, error) {
	taskID, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, err
	}

	var status *entity.TaskStatus
	if input.Status != nil {
		s, err := entity.ParseTaskStatus(string(*input.Status))
		if err != nil {
			return nil, err
		}
		status = &s
	}

	var assigneeIDs []uuid.UUID
	if input.AssigneeIds != nil {
		assigneeIDs = make([]uuid.UUID, 0, len(input.AssigneeIds))
		for _, assigneeID := range input.AssigneeIds {
			id, err := uuid.Parse(assigneeID)
			if err != nil {
				return nil, err
			}
			assigneeIDs = append(assigneeIDs, id)
		}
	}

	taskService := service.NewTaskService(r.DB(ctx))
	task, err := taskService.UpdateTask(ctx, repository.UpdateTaskParams{
		TaskID:      taskID,
		Title:       input.Title,
		Description: input.Description,
		Status:      status,
		IsPrivate:   input.IsPrivate,
		AssigneeIDs: assigneeIDs,
	})
	if err != nil {
		return nil, err
	}

	return &model.UpdateTaskPayload{Task: model.NewTask(*task)}, nil
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
	taskID, err := uuid.Parse(id)
	if err != nil {
		return false, err
	}
	taskService := service.NewTaskService(r.DB(ctx))
	err = taskService.DeleteTask(ctx, taskID)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetTasks is the resolver for the getTasks field.
func (r *queryResolver) GetTasks(ctx context.Context, companyID string) ([]*model.Task, error) {
	cid, err := uuid.Parse(companyID)
	if err != nil {
		return nil, err
	}
	taskRepository := infrastructure.NewTaskRepository(r.DB(ctx))
	tasks, err := taskRepository.GetCompanyTasks(ctx, cid)
	if err != nil {
		return nil, err
	}
	response := make([]*model.Task, 0, len(tasks))
	for _, task := range tasks {
		response = append(response, model.NewTask(*task))
	}
	return response, nil
}

// Company is the resolver for the company field.
func (r *taskResolver) Company(ctx context.Context, obj *model.Task) (*model.Company, error) {
	company, err := storage.GetCompany(ctx, obj.CompanyID)
	if err != nil {
		return nil, err
	}
	return model.NewCompany(*company), nil
}

// Assignees is the resolver for the assignees field.
func (r *taskResolver) Assignees(ctx context.Context, obj *model.Task) ([]*model.User, error) {
	assignees, err := storage.GetAssigneesByTask(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	response := make([]*model.User, 0, len(assignees))
	for _, assignee := range assignees {
		response = append(response, model.NewUser(*assignee))
	}
	return response, nil
}

// Author is the resolver for the author field.
func (r *taskResolver) Author(ctx context.Context, obj *model.Task) (*model.User, error) {
	author, err := storage.GetUser(ctx, obj.AuthorID)
	if err != nil {
		return nil, err
	}
	return model.NewUser(*author), nil
}

// Task returns admin.TaskResolver implementation.
func (r *Resolver) Task() admin.TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }
