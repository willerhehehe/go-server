package service

import (
	"dcswitch/internal/domain"
	"time"
)

type SwitchVersionService struct {
	SwRepo domain.SwitchVersionRepo
}

func (s SwitchVersionService) GetAll() ([]domain.SwitchVersion, error) {
	sw := domain.SwitchVersion{}
	return sw.GetAll(s.SwRepo)
}

func (s SwitchVersionService) Add(name string, time time.Time) error {
	v := domain.SwitchVersion{Name: name, Time: time}
	return v.Add(s.SwRepo)
}

func (s SwitchVersionService) EditName(id int64, name string) (int64, error) {
	return s.SwRepo.EditName(id, name)
}
