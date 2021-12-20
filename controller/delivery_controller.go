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

func (de *DeliveryController) BeforeActivation(ba mvc.BeforeActivation) {
	//通过project_code获取对应的注文列表
	ba.Handle("GET", "/all/{project_code}", "GerAllByProjectCode")
	//通过order_code获取对应的注文列表
	ba.Handle("GET", "/one/{delivery_code}", "GetOneByDeliveryCode")
	//生成纳品书PDF文件
	ba.Handle("GET", "/pdf/{delivery_code}", "DrawDeliveryPdfByDeliveryCode")
	//下载纳品书PDF文件
	ba.Handle("GET", "/download/{destination_name}", "DeliveryPdfDownload")
}

type DeliveryController struct {
	Context         iris.Context
	DeliveryService service.DeliveryService
	Session         *sessions.Session
}

/**
 * url: /v1/delivery/all/{project_code}
 * type：GET
 * descs：通过案件CD获取所有纳品功能
 */
func (de *DeliveryController) GerAllByProjectCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/delivery/all/{project_code} Controller:DeliveryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := de.Context.GetHeader("Authorization")
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

	projectCode := de.Context.Params().Get("project_code")
	delivery := de.DeliveryService.GetDeliveries(projectCode)

	if delivery == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_DELIVERYGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_DELIVERYGET),
			},
		}
	}

	//将查询到的纳品数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range delivery {
		respList = append(respList, item.DeliveryToRespDesc())
	}

	//返回纳品列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_DELIVERYGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_DELIVERYGET),
			"data":    &respList,
		},
	}
}

/**
 * url: /v1/delivery/one/{delivery_code}
 * type：GET
 * descs：通过纳品CD获取某一条纳品信息
 */
func (de *DeliveryController) GetOneByDeliveryCode() mvc.Result {
	const COMMENT = "method:Get url:/v1/delivery/one/{delivery_code} Controller:DeliveryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := de.Context.GetHeader("Authorization")
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

	deliveryCode := de.Context.Params().Get("delivery_code")
	delivery := de.DeliveryService.GetDelivery(deliveryCode)

	if delivery == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_DELIVERYGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_DELIVERYGET),
			},
		}
	}

	//将查询到的纳品数据进行转换成前端需要的内容
	var respList []interface{}
	for _, item := range delivery {
		respList = append(respList, item.DeliveryToRespDesc())
	}

	//返回纳品列表
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_DELIVERYGET,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_DELIVERYGET),
			"data":    &respList,
		},
	}
}

/**
 * 即将添加的注文记录实体
 */
type AddDeliveryEntity struct {
	Id            int64  `json:"id"`
	DeliveryCode  string `json:"delivery_code"`
	EstimateCode  string `json:"estimate_code"`
	ProjectCode   string `json:"project_code"`
	ProjectName   string `json:"project_name"`
	CustomerName  string `json:"customer_name"`
	Deliverables1 string `json:"deliverables1"`
	Deliverables2 string `json:"deliverables2"`
	Deliverables3 string `json:"deliverables3"`
	Quantity1     string `json:"quantity1"`
	Quantity2     string `json:"quantity2"`
	Quantity3     string `json:"quantity3"`
	Memo1         string `json:"memo1"`
	Memo2         string `json:"memo2"`
	Memo3         string `json:"memo3"`
	DeliveryDate  string `json:"delivery_date"`
	Remarks       string `json:"remarks"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
}

/**
 * url: /v1/delivery
 * type：POST
 * descs：添加纳品功能
 */
func (de *DeliveryController) Post() mvc.Result {
	const COMMENT = "method:Post url:/v1/delivery Controller:DeliveryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := de.Context.GetHeader("Authorization")
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

	var deliveryEntity AddDeliveryEntity
	err = de.Context.ReadJSON(&deliveryEntity)
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

	var deliveryInfo model.Delivery

	deliveryInfo.DeliveryCode = deliveryEntity.DeliveryCode
	deliveryInfo.EstimateCode = deliveryEntity.EstimateCode
	deliveryInfo.ProjectCode = deliveryEntity.ProjectCode
	deliveryInfo.ProjectName = deliveryEntity.ProjectName
	deliveryInfo.CustomerName = deliveryEntity.CustomerName
	deliveryInfo.Deliverables1 = deliveryEntity.Deliverables1
	deliveryInfo.Deliverables2 = deliveryEntity.Deliverables2
	deliveryInfo.Deliverables3 = deliveryEntity.Deliverables3
	deliveryInfo.Quantity1 = deliveryEntity.Quantity1
	deliveryInfo.Quantity2 = deliveryEntity.Quantity2
	deliveryInfo.Quantity3 = deliveryEntity.Quantity3
	deliveryInfo.Memo1 = deliveryEntity.Memo1
	deliveryInfo.Memo2 = deliveryEntity.Memo2
	deliveryInfo.Memo3 = deliveryEntity.Memo3
	deliveryInfo.DeliveryDate = deliveryEntity.DeliveryDate
	deliveryInfo.Remarks = deliveryEntity.Remarks
	deliveryInfo.CreatedBy = deliveryEntity.CreatedBy

	isSuccess := de.DeliveryService.SaveDelivery(deliveryInfo)
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
 * url: /v1/delivery
 * type：PUT
 * descs：更新纳品功能
 */
func (de *DeliveryController) Put() mvc.Result {
	const COMMENT = "method:PUT url:/v1/delivery Controller:DeliveryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := de.Context.GetHeader("Authorization")
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

	var deliveryEntity AddDeliveryEntity
	err = de.Context.ReadJSON(&deliveryEntity)
	if err != nil {
		iris.New().Logger().Error(COMMENT + err.Error())
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_DELIVERYUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_DELIVERYUPDATE),
			},
		}
	}

	var deliveryInfo model.Delivery

	deliveryInfo.DeliveryCode = deliveryEntity.DeliveryCode
	deliveryInfo.EstimateCode = deliveryEntity.EstimateCode
	deliveryInfo.ProjectCode = deliveryEntity.ProjectCode
	deliveryInfo.ProjectName = deliveryEntity.ProjectName
	deliveryInfo.CustomerName = deliveryEntity.CustomerName
	deliveryInfo.Deliverables1 = deliveryEntity.Deliverables1
	deliveryInfo.Deliverables2 = deliveryEntity.Deliverables2
	deliveryInfo.Deliverables3 = deliveryEntity.Deliverables3
	deliveryInfo.Quantity1 = deliveryEntity.Quantity1
	deliveryInfo.Quantity2 = deliveryEntity.Quantity2
	deliveryInfo.Quantity3 = deliveryEntity.Quantity3
	deliveryInfo.Memo1 = deliveryEntity.Memo1
	deliveryInfo.Memo2 = deliveryEntity.Memo2
	deliveryInfo.Memo3 = deliveryEntity.Memo3
	deliveryInfo.DeliveryDate = deliveryEntity.DeliveryDate
	deliveryInfo.Remarks = deliveryEntity.Remarks
	deliveryInfo.ModifiedBy = deliveryEntity.ModifiedBy

	isSuccess := de.DeliveryService.UpdateDelivery(deliveryEntity.DeliveryCode, deliveryInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_DELIVERYUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_DELIVERYUPDATE),
			},
		}
	}

	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"type":    utils.RESPMSG_SUCCESS_DELIVERYUPDATE,
			"message": utils.Recode2Text(utils.RESPMSG_SUCCESS_DELIVERYUPDATE),
		},
	}
}

/**
 * url: /v1/delivery/pdf/{delivery_code}
 * type：GET
 * descs：生成纳品书PDF功能
 */
func (de *DeliveryController) DrawDeliveryPdfByDeliveryCode() mvc.Result {
	const COMMENT = "method:Get url: /v1/delivery/pdf/{delivery_code} Controller:DeliveryController" + " "
	iris.New().Logger().Info(COMMENT + "Start")

	token := de.Context.GetHeader("Authorization")
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

	//从前端获取orderCode,并通过orderCode获取order数据
	deliveryCode := de.Context.Params().Get("delivery_code")
	deliveryData := de.DeliveryService.GetDelivery(deliveryCode)

	if deliveryData == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERGET),
			},
		}
	}

	var deliveryDataInfo model.Delivery
	for _, item := range deliveryData {
		deliveryDataInfo = *item
	}
	var fileName string
	if len(deliveryDataInfo.DeliveryPdfNum) != 0 {
		fileName = deliveryDataInfo.DeliveryPdfNum
	} else {
		now := time.Now().Format("2006-01-02")
		_, err = os.Stat(config.InitConfig().FilePath + "/pdf/delivery/" + now)
		if err != nil {
			os.Mkdir(config.InitConfig().FilePath+"/pdf/delivery/"+now, os.ModePerm)
		}

		fileInfo, _ := ioutil.ReadDir(config.InitConfig().FilePath + "/pdf/delivery/" + now)

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

	var deliveryInfo model.Delivery
	deliveryInfo.DeliveryPdfNum = fileName
	isSuccess := de.DeliveryService.UpdateDelivery(deliveryCode, deliveryInfo)
	if !isSuccess {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERUPDATE,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERUPDATE),
			},
		}
	}

	delivery := de.DeliveryService.GetDelivery(deliveryCode)
	if delivery == nil {
		iris.New().Logger().Error(COMMENT + "ERR")
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERGET,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERGET),
			},
		}
	}

	utils.NewDeliveryPdf(delivery)
	//返回pdf文件
	iris.New().Logger().Info(COMMENT + "End")
	return mvc.Response{
		Object: map[string]interface{}{
			"status":   utils.RECODE_OK,
			"type":     utils.RESPMSG_SUCCESS_ESTIMATEGET,
			"message":  utils.Recode2Text(utils.RESPMSG_SUCCESS_ESTIMATEGET),
			"filename": fileName + ".pdf",
		},
	}
}

/**
 * url: /v1/delivery/download/{destination_name}
 * type：GET
 * descs：下载纳品书PDF功能
 */
func (de *DeliveryController) DeliveryPdfDownload() {
	const COMMENT = "method:Get url:/v1/delivery/download/{destination_name} Controller:OrderController" + " "
	iris.New().Logger().Info(COMMENT + "Start")
	destinationName := de.Context.Params().Get("destination_name")
	fileName := config.InitConfig().FilePath + "/pdf/delivery/" + destinationName[0:4] + "-" + destinationName[4:6] + "-" + destinationName[6:8] + "/" + destinationName
	err := de.Context.SendFile(fileName, destinationName)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
	}
}
