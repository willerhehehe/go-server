package service

import "dcswitch/internal/domain"

type SwitchModuleSvc struct {
	SmRepo domain.ModuleRepo
}

func (s SwitchModuleSvc) Start(name string) error {
	md := domain.ModuleDetail{}
	md, err := md.Get(s.SmRepo, name)
	if err != nil {
		return err
	}
	_, err = md.Start(s.SmRepo, md.Id)
	return err
}

func (s SwitchModuleSvc) Success(name string) error {
	md := domain.ModuleDetail{}
	md, err := md.Get(s.SmRepo, name)
	if err != nil {
		return err
	}
	_, err = md.Success(s.SmRepo, md.Id)
	return err
}

func (s SwitchModuleSvc) Fail(id int64, user domain.User, comment string) error {
	md := domain.ModuleDetail{}
	_, err := md.Fail(s.SmRepo, id, user, comment)
	return err
}
