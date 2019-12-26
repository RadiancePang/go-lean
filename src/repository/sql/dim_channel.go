package sql

import "go-learn/components/orm"

//数仓元数据信息
type DimChannelInfo struct {
	OrigChnlLevel1Code string `gorm:"column:orig_chnl_level1_code"`

	OrigChnlLevel1Name string `gorm:"column:orig_chnl_level1_name"`

	OrigChnlLevel2Code string `gorm:"column:orig_chnl_level2_code"`

	OrigChnlLevel2Name string `gorm:"column:orig_chnl_level2_name"`

	ChnlLevel1Code string `gorm:"column:chnl_level1_code"`

	ChnlLevel1Name string `gorm:"column:chnl_level1_name"`

	ChnlLevel2Code string `gorm:"column:chnl_level2_code"`

	ChnlLevel2Name string `gorm:"column:chnl_level2_name"`

	ChnlCode string `gorm:"column:chnl_code"`
}

func (dimChannelInfo *DimChannelInfo) FetchEntity() {

	orm.Db.Where(dimChannelInfo).First(dimChannelInfo)

	return

}
