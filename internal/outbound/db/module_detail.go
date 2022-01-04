package db

import (
	"dcswitch/internal/domain"
)

type ModuleDetailDBRepo struct{}

func (repo ModuleDetailDBRepo) GetModuleDetail(name string) (domain.ModuleDetail, error) {
	// TODO: 未实现
	return domain.ModuleDetail{}, nil
}

func (repo ModuleDetailDBRepo) ModuleDetailStart(detailID int64, detail domain.ModuleDetail) (int64, error) {
	// TODO: 未实现
	return 0, nil
}

func (repo ModuleDetailDBRepo) ModuleDetailSuccess(detailID int64, detail domain.ModuleDetail) (int64, error) {
	// TODO: 未实现
	return 0, nil
}

func (repo ModuleDetailDBRepo) ModuleDetailFail(detailID int64, detail domain.ModuleDetail) (int64, error) {
	// TODO: 未实现
	return 0, nil
}
