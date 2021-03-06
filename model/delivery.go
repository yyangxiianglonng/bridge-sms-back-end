package model

import "time"

/**
 * 纳品信息结构体,用于生成纳品信息表.
 */
type Delivery struct {
	Id             int64     `xorm:"pk unique autoincr" json:"id"`
	DeliveryCode   *string   `xorm:"comment('納品コード')" json:"delivery_code"`
	EstimateCode   *string   `xorm:"comment('見積コード')" json:"estimate_code"`
	ProjectCode    *string   `xorm:"comment('案件コード')" json:"project_code"`
	ProjectName    *string   `xorm:"comment('案件名')" json:"project_name"`
	CustomerName   *string   `xorm:"comment('得意先名')" json:"customer_name"`
	Deliverables1  *string   `xorm:"comment('商品名①')" json:"deliverables1"`
	Deliverables2  *string   `xorm:"comment('商品名②')" json:"deliverables2"`
	Deliverables3  *string   `xorm:"comment('商品名③')" json:"deliverables3"`
	Quantity1      *string   `xorm:"comment('数量①')" json:"quantity1"`
	Quantity2      *string   `xorm:"comment('数量②')" json:"quantity2"`
	Quantity3      *string   `xorm:"comment('数量③')" json:"quantity3"`
	Memo1          *string   `xorm:"comment('備考①')" json:"memo1"`
	Memo2          *string   `xorm:"comment('備考②')" json:"memo2"`
	Memo3          *string   `xorm:"comment('備考③')" json:"memo3"`
	DeliveryDate   *string   `xorm:"comment('納品日')" json:"delivery_date"`
	Remarks        *string   `xorm:"comment('備考')" json:"remarks"`
	DeliveryPdfNum *string   `xorm:"comment('納品書No.')" json:"delivery_pdf_num"`
	CreatedAt      time.Time `xorm:"created comment('作成時間')" json:"created_at"`
	CreatedBy      *string   `xorm:"comment('作成者')" json:"created_by"`
	ModifiedAt     time.Time `xorm:"updated comment('更新時間')" json:"modified_at"`
	ModifiedBy     *string   `xorm:"comment('更新者')" json:"modified_by"`
	DeletedAt      time.Time `xorm:"deleted comment('削除時間')" json:"deleted_at"`
	DeletedBy      *string   `xorm:"comment('削除者')" json:"deleted_by"`
	IsDelete       int64     `json:"is_delete"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (delivery *Delivery) DeliveryToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":               delivery.Id,
		"delivery_code":    delivery.DeliveryCode,
		"estimate_code":    delivery.EstimateCode,
		"project_code":     delivery.ProjectCode,
		"project_name":     delivery.ProjectName,
		"customer_name":    delivery.CustomerName,
		"deliverables1":    delivery.Deliverables1,
		"deliverables2":    delivery.Deliverables2,
		"deliverables3":    delivery.Deliverables3,
		"quantity1":        delivery.Quantity1,
		"quantity2":        delivery.Quantity2,
		"quantity3":        delivery.Quantity3,
		"memo1":            delivery.Memo1,
		"memo2":            delivery.Memo2,
		"memo3":            delivery.Memo3,
		"delivery_date":    delivery.DeliveryDate,
		"remarks":          delivery.Remarks,
		"delivery_pdf_num": delivery.DeliveryPdfNum,
		"created_at":       delivery.CreatedAt,
		"created_by":       delivery.CreatedBy,
		"modified_at":      delivery.ModifiedAt,
		"modified_by":      delivery.ModifiedBy,
		"deleted_at":       delivery.DeletedAt,
		"deleted_by":       delivery.DeletedBy,
		"is_delete":        delivery.IsDelete,
	}
	return
}
