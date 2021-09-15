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

type EstimateController struct {
	Context         iris.Context
	EstimateService service.EstimateService
	Session         *sessions.Session
}

func (es *EstimateController) BeforeActivation(ba mvc.BeforeActivation) {
	//通过project_code获取对应的见积列表
	ba.Handle("GET", "/all/{project_code}", "GetAllByProjectCode")
	//通过estimate_code获取对应的见积情报
	ba.Handle("GET", "/one/{estimate_code}", "GetOneByEstimateCode")
	//通过estimate_code获取对应的见积详细列表
	ba.Handle("GET", "/detail/all/{estimate_code}", "GetAllByEstimateCode")
	//保存见积详细
	ba.Handle("POST", "/detail", "PostEstimateDetail")
	//删除见积详细
	ba.Handle("DELETE", "/detail/{estimate_details_code}", "DeleteDetail")
}

/**
 * url: /v1/estimate/all/{project_code}
 * type：GET
 * descs：获取所有见积功能
 */
func (es *EstimateController) GetAllByProjectCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/estimate/all/{project_code} Controller:EstimateController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := es.Context.GetHeader("Authorization")
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

	projectCode := es.Context.Params().Get("project_code")
	estimate := es.EstimateService.GetEstimates(projectCode)

	if estimate == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEGET),
			},
		}
	}

	//将查询到的见积数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range estimate {
		respList = append(respList, item.EstimateToRespDesc())
	}

	//返回见积列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ESTIMATEGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ESTIMATEGET),
			"data":    &respList,
		},
	}
}

/**
 * url: /v1/estimate/one/{estimate_code}
 * type：GET
 * descs：通过见积CD获取某一条见积信息
 */
func (es *EstimateController) GetOneByEstimateCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/estimate/one/{estimate_code} Controller:EstimateController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := es.Context.GetHeader("Authorization")
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

	estimateCode := es.Context.Params().Get("estimate_code")
	estimate := es.EstimateService.GetEstimate(estimateCode)

	if estimate == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEGET),
			},
		}
	}

	//将查询到的见积数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range estimate {
		respList = append(respList, item.EstimateToRespDesc())
	}
	iris.New().Logger().Info(respList)
	//返回见积列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ESTIMATEGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ESTIMATEGET),
			"data":    &respList,
		},
	}
}

/**
 * 即将添加的见积记录实体
 */
type AddEstimateEntity struct {
	Id                string    `json:"id"`
	EstimateCode      string    `json:"estimate_code"`
	ProjectCode       string    `json:"project_code"`
	ProjectName       string    `json:"project_name"`
	CustomerName      string    `json:"customer_name"`
	EstimateStartDate time.Time `json:"estimate_start_date"`
	EstimateEndDate   time.Time `json:"estimate_end_date"`
	Work1             string    `json:"work_1"`
	Work2             string    `json:"work_2"`
	Work3             string    `json:"work_3"`
	Deliverables1     string    `json:"deliverables_1"`
	Deliverables2     string    `json:"deliverables_2"`
	Deliverables3     string    `json:"deliverables_3"`
	Media1            string    `json:"media_1"`
	Media2            string    `json:"media_2"`
	Media3            string    `json:"media_3"`
	Quantity1         string    `json:"quantity_1"`
	Quantity2         string    `json:"quantity_2"`
	Quantity3         string    `json:"quantity_3"`
	DeliveryDate1     string    `json:"delivery_date_1"`
	DeliveryDate2     string    `json:"delivery_date_2"`
	DeliveryDate3     string    `json:"delivery_date_3"`
	WorkSpace         string    `json:"work_space"`
	SubTotal          string    `json:"sub_total"`
	Tax               string    `json:"tax"`
	Total             string    `json:"total"`
	Supplement        string    `json:"supplement"`
	Remarks           string    `json:"remarks"`
	PaymentConditions string    `json:"payment_conditions"`
	Other             string    `json:"other"`
	IsDelete          int64     `json:"is_delete"`
}

/**
 * url: /v1/estimate
 * type：POST
 * descs：添加见积功能
 */
func (es *EstimateController) Post() mvc.Result {
	const COMMENT = "method:Post url:/v1/estimate Controller:EstimateController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := es.Context.GetHeader("Authorization")
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

	var estimateEntity AddEstimateEntity
	err = es.Context.ReadJSON(&estimateEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEADD),
			},
		}
	}

	var estimateInfo model.Estimate

	estimateInfo.EstimateCode = estimateEntity.EstimateCode
	estimateInfo.ProjectCode = estimateEntity.ProjectCode
	estimateInfo.ProjectName = estimateEntity.ProjectName
	estimateInfo.CustomerName = estimateEntity.CustomerName
	estimateInfo.EstimateStartDate = estimateEntity.EstimateStartDate
	estimateInfo.EstimateEndDate = estimateEntity.EstimateEndDate
	estimateInfo.Work1 = estimateEntity.Work1
	estimateInfo.Work2 = estimateEntity.Work2
	estimateInfo.Work3 = estimateEntity.Work3
	estimateInfo.Deliverables1 = estimateEntity.Deliverables1
	estimateInfo.Deliverables2 = estimateEntity.Deliverables2
	estimateInfo.Deliverables3 = estimateEntity.Deliverables3
	estimateInfo.Media1 = estimateEntity.Media1
	estimateInfo.Media2 = estimateEntity.Media2
	estimateInfo.Media3 = estimateEntity.Media3
	estimateInfo.Quantity1 = estimateEntity.Quantity1
	estimateInfo.Quantity2 = estimateEntity.Quantity2
	estimateInfo.Quantity3 = estimateEntity.Quantity3
	estimateInfo.DeliveryDate1 = estimateEntity.DeliveryDate1
	estimateInfo.DeliveryDate2 = estimateEntity.DeliveryDate2
	estimateInfo.DeliveryDate3 = estimateEntity.DeliveryDate3
	estimateInfo.WorkSpace = estimateEntity.WorkSpace
	estimateInfo.SubTotal = estimateEntity.SubTotal
	estimateInfo.Tax = estimateEntity.Tax
	estimateInfo.Total = estimateEntity.Total
	estimateInfo.Supplement = estimateEntity.Supplement
	estimateInfo.Remarks = estimateEntity.Remarks
	estimateInfo.PaymentConditions = estimateEntity.PaymentConditions
	estimateInfo.Other = estimateEntity.Other
	estimateInfo.IsDelete = estimateEntity.IsDelete

	isSuccess := es.EstimateService.SaveEstimate(estimateInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEADD),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ESTIMATEADD,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ESTIMATEADD),
		},
	}
}

/**
 * url: /v1/estimate
 * type：PUT
 * descs：更新见积功能
 */
func (es *EstimateController) Put() mvc.Result {
	const COMMENT = "method:Put url:/v1/estimate Controller:EstimateController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := es.Context.GetHeader("Authorization")
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

	var estimateEntity AddEstimateEntity
	err = es.Context.ReadJSON(&estimateEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEUPDATE),
			},
		}
	}

	var estimateInfo model.Estimate

	estimateInfo.EstimateCode = estimateEntity.EstimateCode
	estimateInfo.ProjectCode = estimateEntity.ProjectCode
	estimateInfo.ProjectName = estimateEntity.ProjectName
	estimateInfo.CustomerName = estimateEntity.CustomerName
	estimateInfo.EstimateStartDate = estimateEntity.EstimateStartDate
	estimateInfo.EstimateEndDate = estimateEntity.EstimateEndDate
	estimateInfo.Work1 = estimateEntity.Work1
	estimateInfo.Work2 = estimateEntity.Work2
	estimateInfo.Work3 = estimateEntity.Work3
	estimateInfo.Deliverables1 = estimateEntity.Deliverables1
	estimateInfo.Deliverables2 = estimateEntity.Deliverables2
	estimateInfo.Deliverables3 = estimateEntity.Deliverables3
	estimateInfo.Media1 = estimateEntity.Media1
	estimateInfo.Media2 = estimateEntity.Media2
	estimateInfo.Media3 = estimateEntity.Media3
	estimateInfo.Quantity1 = estimateEntity.Quantity1
	estimateInfo.Quantity2 = estimateEntity.Quantity2
	estimateInfo.Quantity3 = estimateEntity.Quantity3
	estimateInfo.DeliveryDate1 = estimateEntity.DeliveryDate1
	estimateInfo.DeliveryDate2 = estimateEntity.DeliveryDate2
	estimateInfo.DeliveryDate3 = estimateEntity.DeliveryDate3
	estimateInfo.WorkSpace = estimateEntity.WorkSpace
	estimateInfo.SubTotal = estimateEntity.SubTotal
	estimateInfo.Tax = estimateEntity.Tax
	estimateInfo.Total = estimateEntity.Total
	estimateInfo.Supplement = estimateEntity.Supplement
	estimateInfo.Remarks = estimateEntity.Remarks
	estimateInfo.PaymentConditions = estimateEntity.PaymentConditions
	estimateInfo.Other = estimateEntity.Other
	estimateInfo.IsDelete = estimateEntity.IsDelete

	isSuccess := es.EstimateService.UpdateEstimate(estimateInfo.EstimateCode, estimateInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEUPDATE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ESTIMATEUPDATE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ESTIMATEUPDATE),
		},
	}
}

/**
 * url: /v1/estimate/detail/all/{estimate_code}
 * type：GET
 * descs：获取所有见积功能
 */
func (es *EstimateController) GetAllByEstimateCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/estimate/detail Controller:EstimateController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := es.Context.GetHeader("Authorization")
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

	estimateCode := es.Context.Params().Get("estimate_code")
	estimateDetail := es.EstimateService.GetEstimateDetails(estimateCode)

	if estimateDetail == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEDETAILGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEDETAILGET),
			},
		}
	}

	//将查询到的见积数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range estimateDetail {
		respList = append(respList, item.EstimateDetailToRespDesc())
	}

	//返回见积列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ESTIMATEDETAILGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ESTIMATEDETAILGET),
			"data":    &respList,
		},
	}
}

/**
 * 即将添加的见积详细记录实体
 */
type AddEstimateDetailEntity struct {
	Id                  int64     `json:"id"`
	EstimateDetailsCode string    `json:"estimate_details_code"`
	EstimateCode        string    `json:"estimate_code"`
	ProductCode         string    `json:"product_code"`
	ProductName         string    `json:"product_name"`
	Quantity            string    `json:"quantity"`
	Price               string    `json:"price"`
	SubTotal            string    `json:"sub_total"`
	Tax                 string    `json:"tax"`
	Total               string    `json:"total"`
	MainFlag            bool      `json:"main_flag"`
	CreatedAt           time.Time `json:"created_at"`
	CreatedBy           string    `json:"created_by"`
	ModifiedAt          time.Time `json:"modified_at"`
	ModifiedBy          string    `json:"modified_by"`
	DeletedAt           time.Time `json:"deleted_at"`
	DeletedBy           string    `json:"deleted_by"`
	IsDelete            int64     `json:"is_delete"`
}

/**
 * url: /v1/estimate/detail
 * type：POST
 * descs：保存见积详细功能
 */
func (es *EstimateController) PostEstimateDetail() mvc.Result {
	const COMMENT = "method:Post url:/v1/estimate/detail Controller:EstimateController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := es.Context.GetHeader("Authorization")
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

	var estimateDetailEntity AddEstimateDetailEntity
	err = es.Context.ReadJSON(&estimateDetailEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEDETAILADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEDETAILADD),
			},
		}
	}

	var estimateDetailInfo model.EstimateDetail

	estimateDetailInfo.EstimateDetailsCode = estimateDetailEntity.EstimateDetailsCode
	estimateDetailInfo.EstimateCode = estimateDetailEntity.EstimateCode
	estimateDetailInfo.ProductCode = estimateDetailEntity.ProductCode
	estimateDetailInfo.ProductName = estimateDetailEntity.ProductName
	estimateDetailInfo.Quantity = estimateDetailEntity.Quantity
	estimateDetailInfo.Price = estimateDetailEntity.Price
	estimateDetailInfo.SubTotal = estimateDetailEntity.SubTotal
	estimateDetailInfo.Tax = estimateDetailEntity.Tax
	estimateDetailInfo.Total = estimateDetailEntity.Total
	estimateDetailInfo.MainFlag = estimateDetailEntity.MainFlag
	estimateDetailInfo.IsDelete = estimateDetailEntity.IsDelete

	isSuccess := es.EstimateService.SaveEstimateDetail(estimateDetailInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEDETAILADD,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEDETAILADD),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ESTIMATEDETAILADD,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ESTIMATEDETAILADD),
		},
	}
}

/**
 * url: /v1/estimate/detail/{estimate_details_code}
 * type：DELETE
 * descs：保存见积详细功能
 */
func (es *EstimateController) DeleteDetail() mvc.Result {
	const COMMENT = "method:Delete url:/v1/estimate/detail/{estimate_details_code} Controller:EstimateController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := es.Context.GetHeader("Authorization")
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

	estimate_details_code := es.Context.Params().Get("estimate_details_code")
	isSuccess := es.EstimateService.DeleteEstimateDetail(estimate_details_code)

	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEDETAILDELETE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEDETAILDELETE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ESTIMATEDETAILDELETE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ESTIMATEDETAILDELETE),
		},
	}
}
