package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/kensei18/enechain-technical-assignment/app/graph/model"
	"github.com/kensei18/enechain-technical-assignment/app/graph/web"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.TaskInput) (*model.CreateTaskPayload, error) {
	panic(fmt.Errorf("not implemented: CreateTask - createTask"))
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, input model.TaskUpdateInput) (*model.UpdateTaskPayload, error) {
	panic(fmt.Errorf("not implemented: UpdateTask - updateTask"))
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteTask - deleteTask"))
}

// GetCompanyTasks is the resolver for the getCompanyTasks field.
func (r *queryResolver) GetCompanyTasks(ctx context.Context) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented: GetCompanyTasks - getCompanyTasks"))
}

// GetUserTasks is the resolver for the getUserTasks field.
func (r *queryResolver) GetUserTasks(ctx context.Context) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented: GetUserTasks - getUserTasks"))
}

// Assignees is the resolver for the assignees field.
func (r *taskResolver) Assignees(ctx context.Context, obj *model.Task) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Assignees - assignees"))
}

// Author is the resolver for the author field.
func (r *taskResolver) Author(ctx context.Context, obj *model.Task) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Author - author"))
}

// Task returns web.TaskResolver implementation.
func (r *Resolver) Task() web.TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }