package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"main/model"
	"main/service"
	"main/utils"
)

type ProjectController struct {
	Context        iris.Context
	ProjectService service.ProjectService
	Session        *sessions.Session
}

func (pr *ProjectController) BeforeActivation(ba mvc.BeforeActivation) {

	//通过project_code获取对应的案件
	ba.Handle("GET", "/one/{project_code}", "GetOneByProjectCode")
}

/**
 * url: /v1/project
 * type：GET
 * descs：获取所有案件功能
 */
func (pr *ProjectController) Get() mvc.Result {
	const COMMENT = "method:Get url:/v1/project Controller:ProjectController" + " "
	iris.New().Logger().Info(COMMENT + "Start")
	projectList := pr.ProjectService.GetProjects()
	if len(projectList) == 0 {
		iris.New().Logger().Info(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERLIST,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USERLIST),
			},
		}
	}

	//将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, project := range projectList {
		respList = append(respList, project.ProjectToRespDesc())
	}

	//返回案件列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_PROJECTGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_PROJECTGET),
			"data":    &respList,
		},
	}
}

/**
 * url: /v1/project/one/{project_code}
 * type：GET
 * descs：通过案件CD案件检索功能
 */
func (pr *ProjectController) GetOneByProjectCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/project/one/{project_code} Controller:ProjectController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	projectCode := pr.Context.Params().Get("project_code")
	project := pr.ProjectService.GetProject(projectCode)

	if project == nil {
		iris.New().Logger().Error(COMMENT + " ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PROJECTGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PROJECTGET),
			},
		}
	}

	//将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range project {
		respList = append(respList, item.ProjectToRespDesc())
	}

	//返回案件列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_PROJECTGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_PROJECTGET),
			"data":    &respList,
		},
	}
}

/**
 * 即将添加的案件记录实体
 */
type AddProjectEntity struct {
	ProjectCode   string `json:"project_code"`
	ProjectName   string `json:"project_name"`
	CustomerCode  string `json:"customer_code"`
	CustomerName  string `json:"customer_name"`
	PersonnelName string `json:"personnel_name"`
	Synopsis      string `json:"synopsis"`
	CreatedBy     string `json:"created_by"`
	IsDelete      int64  `json:"is_delete"`
}

/**
 * url: /v1/project
 * type：POST
 * descs：添加案件功能
 */
func (pr *ProjectController) Post() mvc.Result {
	const COMMENT = "method:Post url:/v1/project Controller:ProjectController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	var projectEntity AddProjectEntity
	err := pr.Context.ReadJSON(&projectEntity)

	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PROJECTADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PROJECTADD),
			},
		}
	}

	var projectInfo model.Project
	projectInfo.ProjectCode = projectEntity.ProjectCode
	projectInfo.ProjectName = projectEntity.ProjectName
	projectInfo.CustomerCode = projectEntity.CustomerCode
	projectInfo.CustomerName = projectEntity.CustomerName
	projectInfo.PersonnelName = projectEntity.PersonnelName
	projectInfo.Synopsis = projectEntity.Synopsis
	projectInfo.CreatedBy = projectEntity.CreatedBy
	projectInfo.IsDelete = projectEntity.IsDelete

	isSuccess := pr.ProjectService.SaveProject(projectInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PROJECTADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PROJECTADD),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_PROJECTADD,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_PROJECTADD),
		},
	}
}

/**
 * url: /v1/project
 * type：PUT
 * descs：更新案件功能
 */
func (pr *ProjectController) Put() mvc.Result {
	const COMMENT = "method:Put url:/v1/project Controller:ProjectController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	var projectEntity AddProjectEntity
	err := pr.Context.ReadJSON(&projectEntity)

	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PROJECTUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PROJECTUPDATE),
			},
		}
	}

	var projectInfo model.Project
	projectInfo.ProjectCode = projectEntity.ProjectCode
	projectInfo.ProjectName = projectEntity.ProjectName
	projectInfo.CustomerCode = projectEntity.CustomerCode
	projectInfo.CustomerName = projectEntity.CustomerName
	projectInfo.PersonnelName = projectEntity.PersonnelName
	projectInfo.Synopsis = projectEntity.Synopsis
	projectInfo.CreatedBy = projectEntity.CreatedBy
	projectInfo.IsDelete = projectEntity.IsDelete

	isSuccess := pr.ProjectService.UpdateProject(projectInfo.ProjectCode, projectInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PROJECTUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_PROJECTUPDATE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_PROJECTUPDATE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_PROJECTUPDATE),
		},
	}
}
