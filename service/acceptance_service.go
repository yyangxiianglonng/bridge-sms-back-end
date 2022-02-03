package service

import (
	"main/model"

	"xorm.io/xorm"
)

/*
* 检收服务接口
 */
type AcceptanceService interface {
	GetAcceptances(projectCode string) []*model.Acceptance
	GetAcceptance(acceptanceCode string) []*model.Acceptance
	SaveAcceptance(acceptance model.Acceptance) bool
	UpdateAcceptance(acceptanceCode string, acceptance model.Acceptance) bool
}

/**
 * 实例化检收服务:服务器
 */
func NewAcceptanceService(engine *xorm.Engine) AcceptanceService {
	return &acceptanceService{
		Engine: engine,
	}
}

/**
 *检收服务实现结构体
 */
type acceptanceService struct {
	Engine *xorm.Engine
}

/*
* 请求某个案件下的所有检收列表数据
 */
func (ac *acceptanceService) GetAcceptances(projectCode string) (acceptanceList []*model.Acceptance) {
	err := ac.Engine.Where("project_code = ?", projectCode).Find(&acceptanceList)

	if err != nil {
		panic(err.Error())
	}
	return
}

/**
 * 通过检收CD获取检收信息
 */
func (ac *acceptanceService) GetAcceptance(acceptanceCode string) (acceptance []*model.Acceptance) {
	err := ac.Engine.Where("acceptance_code = ?", acceptanceCode).Find(&acceptance)

	if err != nil {
		panic(err.Error())
	}
	return
}

/*
 * 保存检收信息
 */
func (ac *acceptanceService) SaveAcceptance(acceptance model.Acceptance) bool {
	_, err := ac.Engine.Insert(&acceptance)
	return err == nil
}

/*
 * 更新检收信息
 */
func (ac *acceptanceService) UpdateAcceptance(acceptanceCode string, acceptance model.Acceptance) bool {
	_, err := ac.Engine.Where("acceptance_code = ?", acceptanceCode).Update(&acceptance)
	return err == nil
}
