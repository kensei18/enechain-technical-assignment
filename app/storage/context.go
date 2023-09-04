package storage

import (
	"context"
	"errors"
)

const dataLoadersKeyName = "dataloadersContextKey"

type dataLoadersKey string

func getLoaders(ctx context.Context) (*Loaders, error) {
	err := errors.New("failed to get dataloaders from context")
	var key dataLoadersKey = dataLoadersKeyName
	value := ctx.Value(key)
	if value == nil {
		return nil, err
	}
	loaders, ok := value.(*Loaders)
	if !ok {
		return nil, err
	}
	return loaders, nil
}

func SetLoaders(ctx context.Context, loaders *Loaders) context.Context {
	var key dataLoadersKey = dataLoadersKeyName
	return context.WithValue(ctx, key, loaders)
}
