package domain

import (
	"github.com/sirupsen/logrus"
	"time"
)

// SwitchModule 切换模块
type SwitchModule struct {
	Version      SwitchVersion
	ProgressRate float64
	ModuleName   string
	BizDomain    string
	Details      []ModuleDetail
	File         string // 切换模块明细附件
}

type ModuleDetail struct {
	Progress
	Id      int64
	Name    string
	Comment string // 失败备注
	User    User   // 失败填写人
}

// Progress 模块明细进度
type Progress struct {
	Status    string // success, fail
	StartTime time.Time
	EndTime   time.Time
}

type ModuleRepo interface {
	GetModuleDetail(name string) (ModuleDetail, error)
	ModuleDetailStart(detailID int64, detail ModuleDetail) (int64, error)
	ModuleDetailSuccess(detailID int64, detail ModuleDetail) (int64, error)
	ModuleDetailFail(detailID int64, detail ModuleDetail) (int64, error)
}

func (m *ModuleDetail) Get(repo ModuleRepo, name string) (ModuleDetail, error) {
	return repo.GetModuleDetail(name)
}

func (m *ModuleDetail) Start(repo ModuleRepo, id int64) (int64, error) {
	m.StartTime = time.Now()
	logrus.Infof("%v start time: %v\n", m.Name, m.StartTime)
	return repo.ModuleDetailStart(id, *m)
}

func (m *ModuleDetail) end() {
	m.EndTime = time.Now()
	logrus.Infof("%v end time: %v\n", m.Name, m.StartTime)
}

func (m *ModuleDetail) Success(repo ModuleRepo, id int64) (int64, error) {
	m.end()
	m.Status = "success"
	return repo.ModuleDetailSuccess(id, *m)
}

func (m *ModuleDetail) Fail(repo ModuleRepo, id int64, user User, comment string) (int64, error) {
	m.end()
	m.User = user
	m.Comment = comment
	m.Status = "fail"
	return repo.ModuleDetailFail(id, *m)
}
