package dict

import (
	"go-learn/common/define"
	"go-learn/components/tools"
	"go-learn/config"
)

type Model struct {
	DictName string `json:"dictName"`

	DictDesc string `json:"DictDesc"`

	DictFeilds []Field `json:"dictFeilds"`
}

type Field struct {
	Code string `json:"code"`

	Name string `json:"name"`

	Remark string `json:"remark"`
}

var dictModels []Model

var dictFieldMap = make(map[string]map[string]Field)

const KeyOther = "OTHER"

func init() {

	config.Load("config/meta/dict_models.json", define.MetaTypeJson, &dictModels)

	parseDictModel()
}

func parseDictModel() {

	for _, dictModel := range dictModels {

		dictFields := dictModel.DictFeilds

		fieldmap := make(map[string]Field)

		for _, dictField := range dictFields {

			fieldmap[dictField.Name] = dictField

		}

		dictFieldMap[dictModel.DictName] = fieldmap

	}

}

func FetchDictInfo(key string, remark string) string {

	var retvalue string

	if fieldmap, ok := dictFieldMap[key]; ok {

		for _, fieldInfo := range fieldmap {

			if fieldInfo.Remark == remark {

				retvalue = fieldInfo.Code

				break

			}

		}

		if tools.IsEmpty(retvalue) {

			if otherFeild, otherOk := fieldmap[KeyOther]; otherOk {

				retvalue = otherFeild.Code

			}

		}

	}

	return retvalue

}

func FetchOtherInfo(key string, name string, remark string) string {

	var retvalue string

	if fieldmap, ok := dictFieldMap[key]; ok {

		for _, fieldInfo := range fieldmap {

			if fieldInfo.Remark == remark && fieldInfo.Name == name {

				retvalue = fieldInfo.Code

				break

			}

		}

		if tools.IsEmpty(retvalue) {

			retvalue = remark

		}

	}

	return retvalue

}
