package repo

import (
	"github.com/1Panel-dev/1Panel/app/model"
	"github.com/1Panel-dev/1Panel/global"
	"gorm.io/gorm"
)

type HostRepo struct{}

type IHostRepo interface {
	Get(opts ...DBOption) (model.Host, error)
	GetList(opts ...DBOption) ([]model.Host, error)
	WithByInfo(info string) DBOption
	Create(host *model.Host) error
	Update(id uint, vars map[string]interface{}) error
	Delete(opts ...DBOption) error
}

func NewIHostService() IHostRepo {
	return &HostRepo{}
}

func (u *HostRepo) Get(opts ...DBOption) (model.Host, error) {
	var host model.Host
	db := global.DB
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.First(&host).Error
	return host, err
}

func (u *HostRepo) GetList(opts ...DBOption) ([]model.Host, error) {
	var hosts []model.Host
	db := global.DB.Model(&model.Host{})
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Find(&hosts).Error
	return hosts, err
}

func (c *HostRepo) WithByInfo(info string) DBOption {
	return func(g *gorm.DB) *gorm.DB {
		if len(info) == 0 {
			return g
		}
		infoStr := "%" + info + "%"
		return g.Where("name LIKE ? OR addr LIKE ?", infoStr, infoStr)
	}
}

func (u *HostRepo) Create(host *model.Host) error {
	return global.DB.Create(host).Error
}

func (u *HostRepo) Update(id uint, vars map[string]interface{}) error {
	return global.DB.Model(&model.Host{}).Where("id = ?", id).Updates(vars).Error
}

func (u *HostRepo) Delete(opts ...DBOption) error {
	db := global.DB
	for _, opt := range opts {
		db = opt(db)
	}
	return db.Delete(&model.Host{}).Error
}