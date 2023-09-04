package storage

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/kensei18/enechain-technical-assignment/app/entity"
)

func GetUser(ctx context.Context, id string) (*entity.User, error) {
	loaders, err := getLoaders(ctx)
	if err != nil {
		return nil, err
	}
	thunk := loaders.UserLoader.Load(ctx, id)
	return thunk()
}

func GetAssigneesByTask(ctx context.Context, taskID string) ([]*entity.User, error) {
	loaders, err := getLoaders(ctx)
	if err != nil {
		return nil, err
	}
	thunk := loaders.TaskAssigneesLoader.Load(ctx, taskID)
	return thunk()
}

func (r *Reader) getUsers(ctx context.Context, keys []string) []*dataloader.Result[*entity.User] {
	output := make([]*dataloader.Result[*entity.User], len(keys))

	users := make([]*entity.User, 0, len(keys))
	if err := r.DB.Find(&users, "id IN ?", keys).Error; err != nil {
		for i := range keys {
			output[i] = &dataloader.Result[*entity.User]{Error: err}
		}
		return output
	}

	userByID := make(map[string]*entity.User, len(users))
	for _, user := range users {
		userByID[user.ID.String()] = user
	}

	for i, key := range keys {
		user, ok := userByID[key]
		if ok {
			output[i] = &dataloader.Result[*entity.User]{Data: user}
		} else {
			output[i] = &dataloader.Result[*entity.User]{Error: fmt.Errorf("user not found: %s", key)}
		}
	}
	return output
}

func (r *Reader) getAssigneesByTask(ctx context.Context, keys []string) []*dataloader.Result[[]*entity.User] {
	output := make([]*dataloader.Result[[]*entity.User], len(keys))

	taskAssignees := make([]*entity.TaskAssignee, 0, len(keys))
	if err := r.DB.Preload("Assignee").Where("task_id IN ?", keys).Find(&taskAssignees).Error; err != nil {
		for i := range keys {
			output[i] = &dataloader.Result[[]*entity.User]{Error: err}
		}
		return output
	}

	assigneeByTaskID := make(map[string][]*entity.User)
	for _, taskAssignee := range taskAssignees {
		taskID := taskAssignee.TaskID.String()
		assigneeByTaskID[taskID] = append(assigneeByTaskID[taskID], &taskAssignee.Assignee)
	}

	for i, key := range keys {
		output[i] = &dataloader.Result[[]*entity.User]{Data: assigneeByTaskID[key]}
	}
	return output
}
