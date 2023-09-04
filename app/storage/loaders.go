package storage

import (
	"context"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/kensei18/enechain-technical-assignment/app/entity"
	"gorm.io/gorm"
)

type Reader struct {
	DB func(context.Context) *gorm.DB
}

type Loaders struct {
	CompanyLoader       *dataloader.Loader[string, *entity.Company]
	UserLoader          *dataloader.Loader[string, *entity.User]
	TaskAssigneesLoader *dataloader.Loader[string, []*entity.User]
}

func NewLoaders(reader *Reader) *Loaders {
	return &Loaders{
		CompanyLoader:       newBatchedLoaderWithoutCache(reader.getCompanies),
		UserLoader:          newBatchedLoaderWithoutCache(reader.getUsers),
		TaskAssigneesLoader: newBatchedLoaderWithoutCache(reader.getAssigneesByTask),
	}
}

func newBatchedLoaderWithoutCache[K comparable, V any](batchFn dataloader.BatchFunc[K, V]) *dataloader.Loader[K, V] {
	return dataloader.NewBatchedLoader[K, V](
		batchFn,
		dataloader.WithCache[K, V](&dataloader.NoCache[K, V]{}),
	)
}
