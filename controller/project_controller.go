package controller

import (
	"main/middleware"
	"main/model"
	"main/service"
	"main/utils"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type ProjectController struct {
	Context        iris.Context
	ProjectService service.ProjectService
	Session        *sessions.Session
}

func (pr *ProjectController) BeforeActivation(ba mvc.BeforeActivation) {
	//通过project_code获取对应的案件
	ba.Handle("GET", "/one/{project_code}", "GetOneByProjectCode", middleware.Author)
	//通过project_code获取对应的时间线数据
	ba.Handle("GET", "/timeline/{project_code}", "GetTimelineByProjectCode")
}

/**
 * url: /v1/project
 * type：GET
 * descs：获取所有案件功能
 */
func (pr *ProjectController) Get() mvc.Result {
	const COMMENT = "method:Get url:/v1/project Controller:ProjectController" + " "
	iris.New().Logger().Info(COMMENT + "Start")
	token := pr.Context.GetHeader("Authorization")
	claim, err := utils.ParseToken(token)

	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	}

	projectList := pr.ProjectService.GetProjects()
	if len(projectList) == 0 {
		iris.New().Logger().Info(COMMENT + "ERR")
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
	token := pr.Context.GetHeader("Authorization")
	claim, err := utils.ParseToken(token)

	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	}

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
 * url: /v1/project/timeline/{project_code}
 * type：GET
 * descs：通过案件CD获取时间线功能
 */
func (pr *ProjectController) GetTimelineByProjectCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/project/timeline/{project_code} Controller:ProjectController" + " "
	iris.New().Logger().Info(COMMENT + "Start")
	token := pr.Context.GetHeader("Authorization")
	claim, err := utils.ParseToken(token)

	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	}

	projectCode := pr.Context.Params().Get("project_code")
	timeline := pr.ProjectService.GetTimeline(projectCode)

	if timeline == nil {
		iris.New().Logger().Error(COMMENT + " ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_TIMELINE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_TIMELINE),
			},
		}
	}

	//将查询到的用户数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range timeline {
		respList = append(respList, item.TimelineToRespDesc())
	}

	//返回案件列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_TIMELINE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_TIMELINE),
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
	ModifiedBy    string `json:"modified_by"`
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
	token := pr.Context.GetHeader("Authorization")
	claim, err := utils.ParseToken(token)

	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	}

	var projectEntity AddProjectEntity
	err = pr.Context.ReadJSON(&projectEntity)

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

	//避免新案件做成画面，连续点击登录更新时造成多次登录案件的问题。
	isExist := pr.ProjectService.ExistProject(projectInfo)
	if isExist {
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
	} else {
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

	token := pr.Context.GetHeader("Authorization")
	claim, err := utils.ParseToken(token)

	if !((err == nil) && (time.Now().Unix() <= claim.ExpiresAt)) {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.RESPMSG_ERROR_SESSION,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SESSION),
			},
		}
	}

	var projectEntity AddProjectEntity
	err = pr.Context.ReadJSON(&projectEntity)

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
	projectInfo.ModifiedBy = projectEntity.ModifiedBy

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
