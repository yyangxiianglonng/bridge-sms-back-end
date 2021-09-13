package model

import (
	"time"
)

/**
 * 案件信息结构体,用于生成案件信息表.
 */
type Project struct {
	Id            int64     `xorm:"pk unique autoincr" json:"id"`
	ProjectCode   string    `json:"project_code"`
	ProjectName   string    `json:"project_name"`
	CustomerCode  string    `json:"customer_code"`
	CustomerName  string    `json:"customer_name"`
	PersonnelName string    `json:"personnel_name"`
	Synopsis      string    `json:"synopsis"`
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
func (project *Project) ProjectToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":             project.Id,
		"project_code":   project.ProjectCode,
		"project_name":   project.ProjectName,
		"customer_code":  project.CustomerCode,
		"customer_name":  project.CustomerName,
		"personnel_name": project.PersonnelName,
		"synopsis":       project.Synopsis,
		"created_at":     project.CreatedAt,
		"created_by":     project.CreatedBy,
		"modified_at":    project.ModifiedAt,
		"modified_by":    project.ModifiedBy,
		"deleted_at":     project.DeletedAt,
		"deleted_by":     project.DeletedBy,
		"is_delete":      project.IsDelete,
	}
	return
}
