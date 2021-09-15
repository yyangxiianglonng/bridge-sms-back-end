package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"main/model"
	"main/service"
	"main/utils"
	"time"
)

func (ac *AcceptanceController) BeforeActivation(ba mvc.BeforeActivation) {
	//通过project_code获取对应的注文列表
	ba.Handle("GET", "/all/{project_code}", "GetAllByProjectCode")
	//通过order_code获取对应的注文列表
	ba.Handle("GET", "/one/{acceptance_code}", "GetOneByAcceptanceCode")
}

type AcceptanceController struct {
	Context           iris.Context
	AcceptanceService service.AcceptanceService
	Session           *sessions.Session
}

/**
 * url: /v1/acceptance/all/{project_code}
 * type：GET
 * descs：通过案件CD获取所有检收功能
 */
func (ac *AcceptanceController) GetAllByProjectCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/acceptance/all/{project_code} Controller:AcceptanceController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := ac.Context.GetHeader("Authorization")
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

	projectCode := ac.Context.Params().Get("project_code")
	acceptance := ac.AcceptanceService.GetAcceptances(projectCode)

	if acceptance == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ACCEPTANCEGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ACCEPTANCEGET),
			},
		}
	}

	//将查询到的检收数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range acceptance {
		respList = append(respList, item.AcceptanceToRespDesc())
	}

	//返回检收列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ACCEPTANCEGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ACCEPTANCEGET),
			"data":    &respList,
		},
	}
}

/**
 * url: /v1/acceptance/one/{acceptance_code}
 * type：GET
 * descs：通过检收CD获取某一条检收信息
 */
func (ac *AcceptanceController) GetOneByAcceptanceCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/acceptance/one/{acceptance_code} Controller:AcceptanceController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := ac.Context.GetHeader("Authorization")
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

	acceptanceCode := ac.Context.Params().Get("acceptance_code")
	acceptance := ac.AcceptanceService.GetAcceptance(acceptanceCode)

	if acceptance == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ACCEPTANCEGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ACCEPTANCEGET),
			},
		}
	}

	//将查询到的检收数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range acceptance {
		respList = append(respList, item.AcceptanceToRespDesc())
	}

	//返回检收列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ACCEPTANCEGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ACCEPTANCEGET),
			"data":    &respList,
		},
	}
}

/**
 * 即将添加的检收记录实体
 */
type AddAcceptanceEntity struct {
	Id             int64  `json:"id"`
	AcceptanceCode string `json:"acceptance_code"`
	DeliveryCode   string `json:"delivery_code"`
	ProjectCode    string `json:"project_code"`
	ProjectName    string `json:"project_name"`
	CustomerName   string `json:"customer_name"`
	Deliverables1  string `json:"deliverables1"`
	Deliverables2  string `json:"deliverables2"`
	Deliverables3  string `json:"deliverables3"`
	Quantity1      string `json:"quantity1"`
	Quantity2      string `json:"quantity2"`
	Quantity3      string `json:"quantity3"`
	Memo1          string `json:"memo1"`
	Memo2          string `json:"memo2"`
	Memo3          string `json:"memo3"`
	AcceptanceDate string `json:"acceptance_date"`
	Remarks        string `json:"remarks"`
	CreatedBy      string `json:"created_by"`
	IsDelete       int64  `json:"is_delete"`
}

/**
 * url: /v1/acceptance
 * type：POST
 * descs：添加检收功能
 */
func (ac *AcceptanceController) Post() mvc.Result {
	const COMMENT = "method:Post url:/v1/acceptance Controller:AcceptanceController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := ac.Context.GetHeader("Authorization")
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

	var acceptanceEntity AddAcceptanceEntity
	err = ac.Context.ReadJSON(&acceptanceEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ACCEPTANCEADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ACCEPTANCEADD),
			},
		}
	}

	var acceptanceInfo model.Acceptance

	acceptanceInfo.AcceptanceCode = acceptanceEntity.AcceptanceCode
	acceptanceInfo.DeliveryCode = acceptanceEntity.DeliveryCode
	acceptanceInfo.ProjectCode = acceptanceEntity.ProjectCode
	acceptanceInfo.ProjectName = acceptanceEntity.ProjectName
	acceptanceInfo.CustomerName = acceptanceEntity.CustomerName
	acceptanceInfo.Deliverables1 = acceptanceEntity.Deliverables1
	acceptanceInfo.Deliverables2 = acceptanceEntity.Deliverables2
	acceptanceInfo.Deliverables3 = acceptanceEntity.Deliverables3
	acceptanceInfo.Quantity1 = acceptanceEntity.Quantity1
	acceptanceInfo.Quantity2 = acceptanceEntity.Quantity2
	acceptanceInfo.Quantity3 = acceptanceEntity.Quantity3
	acceptanceInfo.Memo1 = acceptanceEntity.Memo1
	acceptanceInfo.Memo2 = acceptanceEntity.Memo2
	acceptanceInfo.Memo3 = acceptanceEntity.Memo3
	acceptanceInfo.AcceptanceDate = acceptanceEntity.AcceptanceDate
	acceptanceInfo.Remarks = acceptanceEntity.Remarks
	acceptanceInfo.CreatedBy = acceptanceEntity.CreatedBy
	acceptanceInfo.IsDelete = acceptanceEntity.IsDelete

	isSuccess := ac.AcceptanceService.SaveAcceptance(acceptanceInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ACCEPTANCEADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ACCEPTANCEADD),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ACCEPTANCEADD,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ACCEPTANCEADD),
		},
	}
}

/**
 * url: /v1/acceptance
 * type：PUT
 * descs：更新检收功能
 */
func (ac *AcceptanceController) Put() mvc.Result {
	const COMMENT = "method:Put url:/v1/acceptance Controller:AcceptanceController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := ac.Context.GetHeader("Authorization")
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

	var acceptanceEntity AddAcceptanceEntity
	err = ac.Context.ReadJSON(&acceptanceEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ACCEPTANCEUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ACCEPTANCEUPDATE),
			},
		}
	}

	var acceptanceInfo model.Acceptance

	acceptanceInfo.Id = acceptanceEntity.Id
	acceptanceInfo.AcceptanceCode = acceptanceEntity.AcceptanceCode
	acceptanceInfo.DeliveryCode = acceptanceEntity.DeliveryCode
	acceptanceInfo.ProjectCode = acceptanceEntity.ProjectCode
	acceptanceInfo.ProjectName = acceptanceEntity.ProjectName
	acceptanceInfo.CustomerName = acceptanceEntity.CustomerName
	acceptanceInfo.Deliverables1 = acceptanceEntity.Deliverables1
	acceptanceInfo.Deliverables2 = acceptanceEntity.Deliverables2
	acceptanceInfo.Deliverables3 = acceptanceEntity.Deliverables3
	acceptanceInfo.Quantity1 = acceptanceEntity.Quantity1
	acceptanceInfo.Quantity2 = acceptanceEntity.Quantity2
	acceptanceInfo.Quantity3 = acceptanceEntity.Quantity3
	acceptanceInfo.Memo1 = acceptanceEntity.Memo1
	acceptanceInfo.Memo2 = acceptanceEntity.Memo2
	acceptanceInfo.Memo3 = acceptanceEntity.Memo3
	acceptanceInfo.AcceptanceDate = acceptanceEntity.AcceptanceDate
	acceptanceInfo.Remarks = acceptanceEntity.Remarks
	acceptanceInfo.CreatedBy = acceptanceEntity.CreatedBy
	acceptanceInfo.IsDelete = acceptanceEntity.IsDelete

	isSuccess := ac.AcceptanceService.UpdateAcceptance(acceptanceEntity.AcceptanceCode, acceptanceInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ACCEPTANCEUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ACCEPTANCEUPDATE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ACCEPTANCEUPDATE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ACCEPTANCEUPDATE),
		},
	}
}
