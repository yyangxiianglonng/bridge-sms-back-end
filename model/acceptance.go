package model

import "time"

/**
 * 检收信息结构体,用于生成检收信息表.
 */
type Acceptance struct {
	Id               int64     `xorm:"pk unique autoincr" json:"id"`
	AcceptanceCode   *string   `xorm:"comment('検収コード')" json:"acceptance_code"`
	EstimateCode     *string   `xorm:"comment('見積コード')" json:"estimate_code"`
	DeliveryCode     *string   `xorm:"comment('納品コード')" json:"delivery_code"`
	ProjectCode      *string   `xorm:"comment('案件コード')" json:"project_code"`
	ProjectName      *string   `xorm:"comment('案件名')" json:"project_name"`
	CustomerName     *string   `xorm:"comment('得意先名')" json:"customer_name"`
	Deliverables1    *string   `xorm:"comment('商品名①')" json:"deliverables1"`
	Deliverables2    *string   `xorm:"comment('商品名②')" json:"deliverables2"`
	Deliverables3    *string   `xorm:"comment('商品名③')" json:"deliverables3"`
	Quantity1        *string   `xorm:"comment('数量①')" json:"quantity1"`
	Quantity2        *string   `xorm:"comment('数量②')" json:"quantity2"`
	Quantity3        *string   `xorm:"comment('数量③')" json:"quantity3"`
	Memo1            *string   `xorm:"comment('備考①')" json:"memo1"`
	Memo2            *string   `xorm:"comment('備考②')" json:"memo2"`
	Memo3            *string   `xorm:"comment('備考③')" json:"memo3"`
	AcceptanceDate   *string   `xorm:"comment('検収日')" json:"acceptance_date"`
	Remarks          *string   `xorm:"comment('備考')" json:"remarks"`
	AcceptancePdfNum *string   `xorm:"comment('検収書No.')" json:"acceptance_pdf_num"`
	CreatedAt        time.Time `xorm:"created comment('作成時間')" json:"created_at"`
	CreatedBy        *string   `xorm:"comment('作成者')" json:"created_by"`
	ModifiedAt       time.Time `xorm:"updated comment('更新時間')" json:"modified_at"`
	ModifiedBy       *string   `xorm:"comment('更新者')" json:"modified_by"`
	DeletedAt        time.Time `xorm:"deleted comment('削除時間')" json:"deleted_at"`
	DeletedBy        *string   `xorm:"comment('削除者')" json:"deleted_by"`
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
