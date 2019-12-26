package sql

type FintechBusinessInfo struct {
	OrderId string `gorm:"column:order_id"`

	SysId string `gorm:"column:sys_id"`

	BusinessInfo string `gorm:"column:business_info"`
}

func (FintechBusinessInfo) TableName() string {

	return "fintech_business_info"

}

func (tmp *Template) FetchEntity(report *FintechBusinessInfo) {

	tmp.DB.Where("sys_id = ? and order_id = ?", report.SysId, report.OrderId).Order("update_time desc").First(report)

	return

}
