package usecase

import (
	"context"
	"github.com/asdamwongmantap/api-echo-mongo/crud"
	"github.com/asdamwongmantap/api-echo-mongo/crud/model"
)

type CrudUseCase struct {
	config   *model.EnvConfig
	crudRepo crud.CrudRepositoryI
}

func NewCrudUseCase(config *model.EnvConfig, crudRepo crud.CrudRepositoryI) crud.CrudUseCaseI {
	return &CrudUseCase{
		config:   config,
		crudRepo: crudRepo,
	}
}

func (cuc *CrudUseCase) GetDataUC(ctx context.Context) (resp model.GetDataResponse, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := cuc.crudRepo.GetAllData(ctx)
	if err != nil {
		return resp, err
	}

	return list, err
}

func (cuc *CrudUseCase) InsertDataUC(ctx context.Context, req model.DataProduct) (resp bool, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	err = cuc.crudRepo.InsertData(ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}
