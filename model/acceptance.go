package model

import "time"

/**
 * 检收信息结构体,用于生成检收信息表.
 */
type Acceptance struct {
	Id               int64     `xorm:"pk unique autoincr"json:"id"`
	AcceptanceCode   string    `json:"acceptance_code"`
	EstimateCode     string    `json:"estimate_code"`
	DeliveryCode     string    `json:"delivery_code"`
	ProjectCode      string    `json:"project_code"`
	ProjectName      string    `json:"project_name"`
	CustomerName     string    `json:"customer_name"`
	Deliverables1    string    `json:"deliverables1"`
	Deliverables2    string    `json:"deliverables2"`
	Deliverables3    string    `json:"deliverables3"`
	Quantity1        string    `json:"quantity1"`
	Quantity2        string    `json:"quantity2"`
	Quantity3        string    `json:"quantity3"`
	Memo1            string    `json:"memo1"`
	Memo2            string    `json:"memo2"`
	Memo3            string    `json:"memo3"`
	AcceptanceDate   string    `json:"acceptance_date"`
	Remarks          string    `json:"remarks"`
	AcceptancePdfNum string    `json:"acceptance_pdf_num"`
	CreatedAt        time.Time `xorm:"created" json:"created_at"`
	CreatedBy        string    `json:"created_by"`
	ModifiedAt       time.Time `xorm:"updated" json:"modified_at"`
	ModifiedBy       string    `json:"modified_by"`
	DeletedAt        time.Time `xorm:"deleted" json:"deleted_at"`
	DeletedBy        string    `json:"deleted_by"`
	IsDelete         int64     `json:"is_delete"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (acceptance *Acceptance) AcceptanceToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":                 acceptance.Id,
		"acceptance_code":    acceptance.AcceptanceCode,
		"estimate_code":      acceptance.EstimateCode,
		"delivery_code":      acceptance.DeliveryCode,
		"project_code":       acceptance.ProjectCode,
		"project_name":       acceptance.ProjectName,
		"customer_name":      acceptance.CustomerName,
		"deliverables1":      acceptance.Deliverables1,
		"deliverables2":      acceptance.Deliverables2,
		"deliverables3":      acceptance.Deliverables3,
		"quantity1":          acceptance.Quantity1,
		"quantity2":          acceptance.Quantity2,
		"quantity3":          acceptance.Quantity3,
		"memo1":              acceptance.Memo1,
		"memo2":              acceptance.Memo2,
		"memo3":              acceptance.Memo3,
		"acceptance_date":    acceptance.AcceptanceDate,
		"remarks":            acceptance.Remarks,
		"acceptance_pdf_num": acceptance.AcceptancePdfNum,
		"created_by":         acceptance.CreatedBy,
		"modified_at":        acceptance.ModifiedAt,
		"modified_by":        acceptance.ModifiedBy,
		"deleted_at":         acceptance.DeletedAt,
		"deleted_by":         acceptance.DeletedBy,
		"is_delete":          acceptance.IsDelete,
	}
	return
}
