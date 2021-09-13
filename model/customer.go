package model

import "time"

/**
 * 顾客信息结构体,用于生成客户信息表.
 */
type Customer struct {
	Id            int64     `xorm:"pk unique autoincr" json:"id"`
	CustomerCode  string    `json:"customer_code"`
	CustomerName  string    `json:"customer_name"`
	Department1   string    `json:"department1"`
	Department2   string    `json:"department2"`
	Department3   string    `json:"department3"`
	PersonnelName string    `json:"personnel_name"`
	PostalNumber  string    `json:"postal_number"`
	Address1      string    `json:"address1"`
	Address2      string    `json:"address2"`
	Address3      string    `json:"address3"`
	Telephone     string    `json:"telephone"`
	Fix           string    `json:"fix"`
	CreatedAt     time.Time `xorm:"created" json:"created_at"`
	CreatedBy     string    `json:"created_by"`
	ModifiedAt    time.Time `xorm:"updated" json:"modified_at"`
	ModifiedBy    string    `json:"modified_by"`
	DeletedAt     time.Time `xorm:"deleted" json:"deleted_at"`
	DeletedBy     string    `json:"deleted_by"`
	IsDelete      int64     `json:"is_delete"`
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (customer *Customer) CustomerToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":             customer.Id,
		"customer_code":  customer.CustomerCode,
		"customer_name":  customer.CustomerName,
		"department1":    customer.Department1,
		"department2":    customer.Department2,
		"department3":    customer.Department3,
		"personnel_name": customer.PersonnelName,
		"postal_number":  customer.PostalNumber,
		"address1":       customer.Address1,
		"address2":       customer.Address2,
		"address3":       customer.Address3,
		"telephone":      customer.Telephone,
		"fix":            customer.Fix,
		"created_at":     customer.CreatedAt,
		"created_by":     customer.CreatedBy,
		"modified_at":    customer.ModifiedAt,
		"modified_by":    customer.ModifiedBy,
		"deleted_at":     customer.DeletedAt,
		"deleted_by":     customer.DeletedBy,
		"is_delete":      customer.IsDelete,
	}
	return
}
