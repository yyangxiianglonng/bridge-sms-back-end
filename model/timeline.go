package model

import "time"

type Timeline struct {
	Id          int64     `xorm:"pk autoincr" json:"id"` //主键 用户ID
	ProjectCode string    `json:"project_code"`
	Type        string    `json:"type"`      //创建 更新
	Status      string    `json:"status"`    //案件 见积 注文 纳品 检收 请求
	ChangeAt    time.Time `json:"change_at"` //Status变化的时间
	ChangeBy    string    `json:"change_by"` //变更Status的人
	Changed     string    `json:"changed"`   //变更内容
}

/**
 * 将数据库查询出来的结果进行格式组装成request请求需要的json字段格式
 */
func (timeline *Timeline) TimelineToRespDesc() (respInfo interface{}) {
	respInfo = map[string]interface{}{
		"id":           timeline.Id,
		"project_code": timeline.ProjectCode,
		"type":         timeline.Type,
		"status":       timeline.Status,
		"change_at":    timeline.ChangeAt,
		"change_by":    timeline.ChangeBy,
		"changed":      timeline.Changed,
	}
	return
}
