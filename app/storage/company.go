package storage

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/kensei18/enechain-technical-assignment/app/entity"
)

func GetCompany(ctx context.Context, id string) (*entity.Company, error) {
	loaders, err := getLoaders(ctx)
	if err != nil {
		return nil, err
	}
	thunk := loaders.CompanyLoader.Load(ctx, id)
	return thunk()
}

func (r *Reader) getCompanies(ctx context.Context, keys []string) []*dataloader.Result[*entity.Company] {
	output := make([]*dataloader.Result[*entity.Company], len(keys))

	companies := make([]*entity.Company, 0, len(keys))
	if err := r.DB.Find(&companies, "id IN ?", keys).Error; err != nil {
		for i := range keys {
			output[i] = &dataloader.Result[*entity.Company]{Error: err}
		}
		return output
	}

	companyByID := make(map[string]*entity.Company, len(companies))
	for _, company := range companies {
		companyByID[company.ID.String()] = company
	}

	for i, key := range keys {
		company, ok := companyByID[key]
		if ok {
			output[i] = &dataloader.Result[*entity.Company]{Data: company}
		} else {
			output[i] = &dataloader.Result[*entity.Company]{Error: fmt.Errorf("company not found: %s", key)}
		}
	}
	return output
}
