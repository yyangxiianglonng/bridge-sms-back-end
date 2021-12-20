package model

import "time"

/**
 * 纳品信息结构体,用于生成纳品信息表.
 */
type Delivery struct {
	Id             int64     `xorm:"pk unique autoincr"json:"id"`
	DeliveryCode   string    `json:"delivery_code"`
	EstimateCode   string    `json:"estimate_code"`
	ProjectCode    string    `json:"project_code"`
	ProjectName    string    `json:"project_name"`
	CustomerName   string    `json:"customer_name"`
	Deliverables1  string    `json:"deliverables1"`
	Deliverables2  string    `json:"deliverables2"`
	Deliverables3  string    `json:"deliverables3"`
	Quantity1      string    `json:"quantity1"`
	Quantity2      string    `json:"quantity2"`
	Quantity3      string    `json:"quantity3"`
	Memo1          string    `json:"memo1"`
	Memo2          string    `json:"memo2"`
	Memo3          string    `json:"memo3"`
	DeliveryDate   string    `json:"delivery_date"`
	Remarks        string    `json:"remarks"`
	DeliveryPdfNum string    `json:"delivery_pdf_num"`
	CreatedAt      time.Time `xorm:"created" json:"created_at"`
	CreatedBy      string    `json:"created_by"`
	ModifiedAt     time.Time `xorm:"updated" json:"modified_at"`
	ModifiedBy     string    `json:"modified_by"`
	DeletedAt      time.Time `xorm:"deleted" json:"deleted_at"`
	DeletedBy      string    `json:"deleted_by"`
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
