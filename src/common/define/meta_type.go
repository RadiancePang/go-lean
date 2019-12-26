package define

import "strings"

/*元数据定义*/
const (
	MetaTypeJson    = "json"
	MetaTypeYaml    = "yaml"
	MetaTypeInt     = "integer"
	MetaTypeDecimal = "bigdecimal"
	MetaTypeString  = "string"
	ModelYxScore    = "yxscore-features"
	DefaultFormula  = "cal_iter_default"
)

var initValueMap = map[string]string{
	MetaTypeInt:     "0",
	MetaTypeDecimal: "0.00",
	MetaTypeString:  "",
	ModelYxScore:    "-999",
}

/*获取初始化值*/
func FetchInitValue(code string) string {

	keycode := strings.ToLower(code)

	return initValueMap[keycode]

}
