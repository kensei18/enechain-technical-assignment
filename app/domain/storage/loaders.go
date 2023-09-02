package storage

import (
	"github.com/graph-gophers/dataloader/v7"
	"gorm.io/gorm"
)

type Reader struct {
	DB *gorm.DB
}

type Loaders struct{}

func NewLoaders(reader *Reader) *Loaders {
	return &Loaders{}
}

func newBatchedLoaderWithoutCache[K comparable, V any](batchFn dataloader.BatchFunc[K, V]) *dataloader.Loader[K, V] {
	return dataloader.NewBatchedLoader[K, V](
		batchFn,
		dataloader.WithCache[K, V](&dataloader.NoCache[K, V]{}),
	)
}
