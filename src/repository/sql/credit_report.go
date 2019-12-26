package sql

import "go-learn/components/orm"

type CreditReport struct {
	OrderId string `gorm:"column:order_id"`

	SysId string `gorm:"column:sys_id"`

	Type string `gorm:"column:type"`

	JsonContext string `gorm:"column:json_context"`
}

func (CreditReport) TableName() string {

	return "credit_dfs_fileid"

}

func (report *CreditReport) FetchEntity() {

	orm.Db.Where("type = ? and sys_id = ? and order_id = ?", report.Type, report.SysId, report.OrderId).Order("update_time desc").First(report)

	return

}
