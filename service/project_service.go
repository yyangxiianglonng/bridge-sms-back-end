package service

import (
	"main/model"
	"xorm.io/xorm"
)

/**
*	案件服务接口
 */
type ProjectService interface {
	GetProjects() []*model.Project
	GetProject(projectCode string) []*model.Project
	GetTimeline(projectCode string) (timeline []*model.Timeline)
	SaveProject(project model.Project) bool
	UpdateProject(projectCode string, estimate model.Project) bool
}

/**
 * 实例化案件服务:服务器
 */
func NewProjectService(engine *xorm.Engine) ProjectService {
	return &projectService{
		Engine: engine,
	}
}

/**
 * 案件服务实现结构体
 */
type projectService struct {
	Engine *xorm.Engine
}

/**
 * 请求案件列表数据
 */
func (pr *projectService) GetProjects() (projectList []*model.Project) {
	err := pr.Engine.Where("is_delete = ?", 0).Find(&projectList)

	//debug用↓
	//for _, project := range projectList {
	//	iris.New().Logger().Info(*project)
	//}
	//debug用↑
	if err != nil {
		panic(err.Error())
	}
	return
}

/**
 * 通过案件CD获取单个案件信息服务
 */
func (pr *projectService) GetProject(projectCode string) (project []*model.Project) {
	err := pr.Engine.Where("is_delete = ?", 0).And("project_code = ?", projectCode).Find(&project)
	if err != nil {
		panic(err.Error())
	}
	return
}

/**
 * 保存案件服务
 */
func (pr *projectService) SaveProject(project model.Project) bool {
	_, err := pr.Engine.Insert(&project)
	return err == nil
}

/**
 * 更新案件信息服务
 */
func (pr *projectService) UpdateProject(projectCode string, project model.Project) bool {
	_, err := pr.Engine.Where("project_code = ?", projectCode).Update(project)
	return err == nil
}

/**
 * 根据案件CD获取时间线数据服务
 */
func (pr *projectService) GetTimeline(projectCode string) (timeline []*model.Timeline) {
	err := pr.Engine.Where("project_code = ?", projectCode).Find(&timeline)

	if err != nil {
		panic(err.Error())
	}
	return
}
