package service

import (
	"github.com/kataras/iris/v12"
	"main/model"
	"xorm.io/xorm"
)

/**
*	见积服务接口
 */
type EstimateService interface {
	//见积头服务接口
	GetEstimates(project_code string) []*model.Estimate
	GetEstimate(estimate_code string) []*model.Estimate
	SaveEstimate(estimate model.Estimate) bool
	UpdateEstimate(estimateCode string, estimate model.Estimate) bool

	//见积详细服务接口
	GetEstimateDetails(estimate_code string) []*model.EstimateDetail
	SaveEstimateDetail(estimate model.EstimateDetail) bool
	UpdateEstimateDetail(estimateDetailsCode string, estimate model.EstimateDetail) bool
	DeleteEstimateDetail(estimateDetailsCode string, estimateDetail model.EstimateDetail) bool
}

/**
*	实例化见积服务:服务器
 */
func NewEstimateService(engine *xorm.Engine) EstimateService {
	return &estimateService{
		Engine: engine,
	}
}

/**
*	见积服务实现结构体
 */
type estimateService struct {
	Engine *xorm.Engine
}

/**
 * 请求某个案件下的所有见积列表数据
 */
func (es *estimateService) GetEstimates(projectCode string) (estimateList []*model.Estimate) {
	err := es.Engine.Where("is_delete = ?", 0).And("project_code = ?", projectCode).Find(&estimateList)

	if err != nil {
		panic(err.Error())
	}
	return
}

/**
 * 通过见积CD获取见积信息
 */
func (es *estimateService) GetEstimate(estimateCode string) (estimate []*model.Estimate) {
	err := es.Engine.Where("is_delete = ?", 0).And("estimate_code = ?", estimateCode).Find(&estimate)

	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}
	return
}

/**
 * 保存见积服务
 */
func (es *estimateService) SaveEstimate(estimate model.Estimate) bool {
	_, err := es.Engine.Insert(&estimate)
	return err == nil
}

/**
 * 更新见积服务
 */
func (es *estimateService) UpdateEstimate(estimateCode string, estimate model.Estimate) bool {
	_, err := es.Engine.Where("estimate_code = ?", estimateCode).Update(&estimate)
	return err == nil
}

/**
 * 获取某个见积下面所有见积详细列表服务
 */
func (es *estimateService) GetEstimateDetails(estimateCode string) (estimateDetailList []*model.EstimateDetail) {
	err := es.Engine.Where("is_delete = ?", 0).And("estimate_code = ?", estimateCode).Find(&estimateDetailList)

	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}
	return
}

/**
 * 保存见积详细服务
 */
func (es *estimateService) SaveEstimateDetail(estimateDetail model.EstimateDetail) bool {
	_, err := es.Engine.Insert(&estimateDetail)
	return err == nil
}

/**
 * 更新见积详细服务
 */
func (es *estimateService) UpdateEstimateDetail(estimateDetailsCode string, estimateDetail model.EstimateDetail) bool {
	_, err := es.Engine.Where("estimate_details_code = ?", estimateDetailsCode).Update(&estimateDetail)
	return err == nil
}

/**
 * 删除见积详细服务
 */
func (es *estimateService) DeleteEstimateDetail(estimateDetailsCode string, estimateDetail model.EstimateDetail) bool {
	//var estimateDetail model.EstimateDetail
	_, err := es.Engine.Where("estimate_details_code = ?", estimateDetailsCode).Update(&estimateDetail)
	_, err = es.Engine.Where("estimate_details_code = ?", estimateDetailsCode).Delete(&estimateDetail)
	return err == nil
}
