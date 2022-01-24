package controller

import (
	"io/ioutil"
	"main/config"
	"main/model"
	"main/service"
	"main/utils"
	"os"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

// const FILEPATH = "/Users/yangxianglong/go/Vue_Iris/back-end/static/file/pdf/estimate/"

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
	//更新见积详细
	ba.Handle("PUT", "/detail", "PutEstimateDetail")
	//删除见积详细
	ba.Handle("DELETE", "/detail/{estimate_details_code}", "DeleteDetail")
	//生成见积书PDF文件
	ba.Handle("GET", "/pdf/{estimate_code}", "DrawPdfByEstimateCode")
	//下载见积书PDF文件
	ba.Handle("GET", "/download/{destination_name}", "PdfDownload")
}

/**
 * url: /v1/estimate
 * type：GET
 * descs：获取所有见积功能
 */
func (es *EstimateController) Get() mvc.Result {
	const COMMENT = "method:Get url:/v1/estimate Controller:EstimateController" + " "
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

	estimate := es.EstimateService.GetEstimateAll()
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
 * url: /v1/estimate/search
 * type：GET
 * descs：获取所有见积功能
 */
func (es *EstimateController) GetSearch() mvc.Result {
	const COMMENT = "method:Get url:/v1/estimate/search Controller:EstimateController" + " "
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

	params := es.Context.URLParams()
	ProjectCode := params["project_code"]
	ProjectName := params["project_name"]
	CustomerName := params["customer_name"]
	var estimateInfo model.Estimate
	estimateInfo.ProjectCode = &ProjectCode
	estimateInfo.ProjectName = &ProjectName
	// estimateInfo.CustomerCode = params["customer_code"]
	estimateInfo.CustomerName = &CustomerName

	estimate := es.EstimateService.SearchEstimates(estimateInfo)
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

	//返回见积
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
	EstimateName      string    `json:"estimate_name"`
	ProjectCode       string    `json:"project_code"`
	ProjectName       string    `json:"project_name"`
	CustomerName      string    `json:"customer_name"`
	EstimateStartDate time.Time `json:"estimate_start_date"`
	EstimateEndDate   time.Time `json:"estimate_end_date"`
	Work1             string    `json:"work1"`
	Work2             string    `json:"work2"`
	Work3             string    `json:"work3"`
	Deliverables1     string    `json:"deliverables1"`
	Deliverables2     string    `json:"deliverables2"`
	Deliverables3     string    `json:"deliverables3"`
	Media1            string    `json:"media1"`
	Media2            string    `json:"media2"`
	Media3            string    `json:"media3"`
	Quantity1         string    `json:"quantity1"`
	Quantity2         string    `json:"quantity2"`
	Quantity3         string    `json:"quantity3"`
	DeliveryDate1     string    `json:"delivery_date1"`
	DeliveryDate2     string    `json:"delivery_date2"`
	DeliveryDate3     string    `json:"delivery_date3"`
	WorkSpace         string    `json:"work_space"`
	SubTotal          string    `json:"sub_total"`
	Tax               string    `json:"tax"`
	Total             string    `json:"total"`
	Supplement        string    `json:"supplement"`
	Remarks           string    `json:"remarks"`
	PaymentConditions string    `json:"payment_conditions"`
	Other             string    `json:"other"`
	CreatedBy         string    `json:"created_by"`
	ModifiedBy        string    `json:"modified_by"`
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

	estimateInfo.EstimateCode = &estimateEntity.EstimateCode
	estimateInfo.EstimateName = &estimateEntity.EstimateName
	estimateInfo.ProjectCode = &estimateEntity.ProjectCode
	estimateInfo.ProjectName = &estimateEntity.ProjectName
	estimateInfo.CustomerName = &estimateEntity.CustomerName
	estimateInfo.EstimateStartDate = estimateEntity.EstimateStartDate
	estimateInfo.EstimateEndDate = estimateEntity.EstimateEndDate
	estimateInfo.Work1 = &estimateEntity.Work1
	estimateInfo.Work2 = &estimateEntity.Work2
	estimateInfo.Work3 = &estimateEntity.Work3
	estimateInfo.Deliverables1 = &estimateEntity.Deliverables1
	estimateInfo.Deliverables2 = &estimateEntity.Deliverables2
	estimateInfo.Deliverables3 = &estimateEntity.Deliverables3
	estimateInfo.Media1 = &estimateEntity.Media1
	estimateInfo.Media2 = &estimateEntity.Media2
	estimateInfo.Media3 = &estimateEntity.Media3
	estimateInfo.Quantity1 = &estimateEntity.Quantity1
	estimateInfo.Quantity2 = &estimateEntity.Quantity2
	estimateInfo.Quantity3 = &estimateEntity.Quantity3
	estimateInfo.DeliveryDate1 = &estimateEntity.DeliveryDate1
	estimateInfo.DeliveryDate2 = &estimateEntity.DeliveryDate2
	estimateInfo.DeliveryDate3 = &estimateEntity.DeliveryDate3
	estimateInfo.WorkSpace = &estimateEntity.WorkSpace
	estimateInfo.SubTotal = &estimateEntity.SubTotal
	estimateInfo.Tax = &estimateEntity.Tax
	estimateInfo.Total = &estimateEntity.Total
	estimateInfo.Supplement = &estimateEntity.Supplement
	estimateInfo.Remarks = &estimateEntity.Remarks
	estimateInfo.PaymentConditions = &estimateEntity.PaymentConditions
	estimateInfo.Other = &estimateEntity.Other
	estimateInfo.CreatedBy = &estimateEntity.CreatedBy
	// estimateInfo.IsDelete = estimateEntity.IsDelete

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

	estimateInfo.EstimateCode = &estimateEntity.EstimateCode
	estimateInfo.EstimateName = &estimateEntity.EstimateName
	estimateInfo.ProjectCode = &estimateEntity.ProjectCode
	estimateInfo.ProjectName = &estimateEntity.ProjectName
	estimateInfo.CustomerName = &estimateEntity.CustomerName
	estimateInfo.EstimateStartDate = estimateEntity.EstimateStartDate
	estimateInfo.EstimateEndDate = estimateEntity.EstimateEndDate
	estimateInfo.Work1 = &estimateEntity.Work1
	estimateInfo.Work2 = &estimateEntity.Work2
	estimateInfo.Work3 = &estimateEntity.Work3
	estimateInfo.Deliverables1 = &estimateEntity.Deliverables1
	estimateInfo.Deliverables2 = &estimateEntity.Deliverables2
	estimateInfo.Deliverables3 = &estimateEntity.Deliverables3
	estimateInfo.Media1 = &estimateEntity.Media1
	estimateInfo.Media2 = &estimateEntity.Media2
	estimateInfo.Media3 = &estimateEntity.Media3
	estimateInfo.Quantity1 = &estimateEntity.Quantity1
	estimateInfo.Quantity2 = &estimateEntity.Quantity2
	estimateInfo.Quantity3 = &estimateEntity.Quantity3
	estimateInfo.DeliveryDate1 = &estimateEntity.DeliveryDate1
	estimateInfo.DeliveryDate2 = &estimateEntity.DeliveryDate2
	estimateInfo.DeliveryDate3 = &estimateEntity.DeliveryDate3
	estimateInfo.WorkSpace = &estimateEntity.WorkSpace
	estimateInfo.SubTotal = &estimateEntity.SubTotal
	estimateInfo.Tax = &estimateEntity.Tax
	estimateInfo.Total = &estimateEntity.Total
	estimateInfo.Supplement = &estimateEntity.Supplement
	estimateInfo.Remarks = &estimateEntity.Remarks
	estimateInfo.PaymentConditions = &estimateEntity.PaymentConditions
	estimateInfo.Other = &estimateEntity.Other
	estimateInfo.ModifiedBy = &estimateEntity.ModifiedBy
	// estimateInfo.IsDelete = estimateEntity.IsDelete

	isSuccess := es.EstimateService.UpdateEstimate(*estimateInfo.EstimateCode, estimateInfo)
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
	Index               string    `json:"index"`
	EstimateDetailsCode string    `json:"estimate_details_code"`
	EstimateCode        string    `json:"estimate_code"`
	ProductCode         string    `json:"product_code"`
	ProductName         string    `json:"product_name"`
	Quantity            string    `json:"quantity"`
	Price               string    `json:"price"`
	SubTotal            string    `json:"sub_total"`
	Tax                 string    `json:"tax"`
	Total               string    `json:"total"`
	Remarks             string    `json:"remarks"`
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

	estimateDetailInfo.Index = &estimateDetailEntity.Index
	estimateDetailInfo.EstimateDetailsCode = &estimateDetailEntity.EstimateDetailsCode
	estimateDetailInfo.EstimateCode = &estimateDetailEntity.EstimateCode
	estimateDetailInfo.ProductCode = &estimateDetailEntity.ProductCode
	estimateDetailInfo.ProductName = &estimateDetailEntity.ProductName
	estimateDetailInfo.Quantity = &estimateDetailEntity.Quantity
	estimateDetailInfo.Price = &estimateDetailEntity.Price
	estimateDetailInfo.SubTotal = &estimateDetailEntity.SubTotal
	estimateDetailInfo.Tax = &estimateDetailEntity.Tax
	estimateDetailInfo.Total = &estimateDetailEntity.Total
	estimateDetailInfo.Remarks = &estimateDetailEntity.Remarks
	estimateDetailInfo.MainFlag = &estimateDetailEntity.MainFlag
	estimateDetailInfo.CreatedBy = &estimateDetailEntity.CreatedBy
	// estimateDetailInfo.IsDelete = &estimateDetailEntity.IsDelete

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
 * url: /v1/estimate/detail
 * type：Put
 * descs：更新见积详细功能
 */
func (es *EstimateController) PutEstimateDetail() mvc.Result {
	const COMMENT = "method:Put url:/v1/estimate/detail Controller:EstimateController" + " "
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
				"type":    utils.RESPMSG_ERROR_ESTIMATEDETAILUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEDETAILUPDATE),
			},
		}
	}

	var estimateDetailInfo model.EstimateDetail
	estimateDetailInfo.Index = &estimateDetailEntity.Index
	estimateDetailInfo.EstimateDetailsCode = &estimateDetailEntity.EstimateDetailsCode
	estimateDetailInfo.EstimateCode = &estimateDetailEntity.EstimateCode
	estimateDetailInfo.ProductCode = &estimateDetailEntity.ProductCode
	estimateDetailInfo.ProductName = &estimateDetailEntity.ProductName
	estimateDetailInfo.Quantity = &estimateDetailEntity.Quantity
	estimateDetailInfo.Price = &estimateDetailEntity.Price
	estimateDetailInfo.SubTotal = &estimateDetailEntity.SubTotal
	estimateDetailInfo.Tax = &estimateDetailEntity.Tax
	estimateDetailInfo.Total = &estimateDetailEntity.Total
	estimateDetailInfo.Remarks = &estimateDetailEntity.Remarks
	estimateDetailInfo.MainFlag = &estimateDetailEntity.MainFlag
	estimateDetailInfo.ModifiedBy = &estimateDetailEntity.ModifiedBy
	// estimateDetailInfo.IsDelete = &estimateDetailEntity.IsDelete

	isSuccess := es.EstimateService.UpdateEstimateDetail(estimateDetailEntity.EstimateDetailsCode, estimateDetailInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEDETAILUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEDETAILUPDATE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_ESTIMATEDETAILUPDATE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_ESTIMATEDETAILUPDATE),
		},
	}
}

/**
 * url: /v1/estimate/pdf/{estimate_code}
 * type：GET
 * descs：生成见积书PDF功能
 */
func (es *EstimateController) DrawPdfByEstimateCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/estimate/pdf/{estimate_code} Controller:EstimateController" + " "
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
	//从前端获取estimateCode,并通过estimateCode获取estimate数据
	estimateCode := es.Context.Params().Get("estimate_code")
	estimateData := es.EstimateService.GetEstimate(estimateCode)

	if estimateData == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEGET),
			},
		}
	}

	var estimateDataInfo model.Estimate
	for _, item := range estimateData {
		estimateDataInfo = *item
	}

	var fileName string

	if estimateDataInfo.EstimatePdfNum != nil {
		fileName = *estimateDataInfo.EstimatePdfNum
	} else {
		now := time.Now().Format("2006-01-02")
		_, err = os.Stat(config.InitConfig().FilePath + "/pdf/estimate/" + now)
		if err != nil {
			os.Mkdir(config.InitConfig().FilePath+"/pdf/estimate/"+now, os.ModePerm)
		}

		fileInfo, _ := ioutil.ReadDir(config.InitConfig().FilePath + "/pdf/estimate/" + now)

		var files []string
		for _, file := range fileInfo {
			files = append(files, file.Name())
		}

		if len(files) < 10 {
			fileName = time.Now().Format("20060102") + "0" + strconv.Itoa(len(files)+1)
		} else {
			fileName = time.Now().Format("20060102") + strconv.Itoa(len(files)+1)
		}
	}

	var estimateInfo model.Estimate
	estimateInfo.EstimatePdfNum = &fileName
	isSuccess := es.EstimateService.UpdateEstimate(estimateCode, estimateInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEDETAILUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEDETAILUPDATE),
			},
		}
	}

	estimate := es.EstimateService.GetEstimate(estimateCode)
	estimateDetail := es.EstimateService.GetEstimateDetails(estimateCode)

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
	// var respList []interface{}
	// for _, item := range estimate {
	// 	respList = append(respList, item.EstimateToRespDesc())
	// }

	utils.NewEstimatePdf(estimate, estimateDetail)
	//返回pdf文件
	iris.New().Logger().Info(COMMENT + "End")

	return mvc.Response{
		Object: map[string]interface{}{
			"status":   utils.RECODE_OK,
			"type":     utils.RESPMSG_SUCCESS_ESTIMATEGET,
			"message":  utils.Recode2Text(utils.RESPMSG_SUCCESS_ESTIMATEGET),
			"filename": fileName + "_見積書_" + *estimateDataInfo.CustomerName + "様_" + *estimateDataInfo.EstimateName + ".pdf",
		},
	}
}

/**
 * url: /v1/estimate/download/{destination_name}
 * type：GET
 * descs：下载见积书PDF功能
 */
func (es *EstimateController) PdfDownload() {
	const COMMENT = "method:Get url:/v1/estimate/download/{destination_name} Controller:EstimateController" + " "
	iris.New().Logger().Info(COMMENT + "Start")
	destinationName := es.Context.Params().Get("destination_name")
	// fileName := config.InitConfig().FilePath + "/pdf/estimate/" + time.Now().Format("2006-01-02") + "/" + destinationName
	fileName := config.InitConfig().FilePath + "/pdf/estimate/" + destinationName[0:4] + "-" + destinationName[4:6] + "-" + destinationName[6:8] + "/" + destinationName
	err := es.Context.SendFile(fileName, destinationName)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}
}

/**
 * url: /v1/estimate/detail/{estimate_details_code}
 * type：DELETE
 * descs：删除见积详细功能
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

	var estimateDetailEntity AddEstimateDetailEntity
	err = es.Context.ReadJSON(&estimateDetailEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ESTIMATEDETAILUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ESTIMATEDETAILUPDATE),
			},
		}
	}

	var estimateDetailInfo model.EstimateDetail

	estimateDetailInfo.Index = &estimateDetailEntity.Index
	estimateDetailInfo.EstimateDetailsCode = &estimateDetailEntity.EstimateDetailsCode
	estimateDetailInfo.EstimateCode = &estimateDetailEntity.EstimateCode
	estimateDetailInfo.ProductCode = &estimateDetailEntity.ProductCode
	estimateDetailInfo.ProductName = &estimateDetailEntity.ProductName
	estimateDetailInfo.Quantity = &estimateDetailEntity.Quantity
	estimateDetailInfo.Price = &estimateDetailEntity.Price
	estimateDetailInfo.SubTotal = &estimateDetailEntity.SubTotal
	estimateDetailInfo.Tax = &estimateDetailEntity.Tax
	estimateDetailInfo.Total = &estimateDetailEntity.Total
	estimateDetailInfo.MainFlag = &estimateDetailEntity.MainFlag
	estimateDetailInfo.ModifiedBy = &estimateDetailEntity.ModifiedBy
	estimateDetailInfo.DeletedBy = &estimateDetailEntity.DeletedBy
	// estimateDetailInfo.IsDelete = &estimateDetailEntity.IsDelete

	estimate_details_code := es.Context.Params().Get("estimate_details_code")
	isSuccess := es.EstimateService.DeleteEstimateDetail(estimate_details_code, estimateDetailInfo)

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
