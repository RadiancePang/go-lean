package sql

import "go-learn/components/orm"

type DimCarStyleExt struct {
	MakeId string `gorm:"column:make_id"`

	MakeName string `gorm:"column:make_name"`

	StyleId string `gorm:"column:style_id"`

	StyleName string `gorm:"column:style_name"`

	SeriesLevelId string `gorm:"column:series_level_id"`

	SeriesLevelName string `gorm:"column:series_level_name"`
}

func (dimCarStyleExt *DimCarStyleExt) FetchEntity() {

	orm.Db.Where(dimCarStyleExt).First(dimCarStyleExt)

	return

}
