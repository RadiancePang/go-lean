package provider

import (
	"reflect"
)

func (provider FormulaProvier) DictModel() string {

	featureInfo := provider.FeatureInfo

	dataMap := provider.DataMap

	dictValue := provider.CreditDefaultStr(featureInfo.Type, featureInfo.DataSetId)

	var fetchValue string

	fetechValueObject := tools.IterMap(featureInfo.Name, dataMap)

	valueKind := reflect.ValueOf(fetechValueObject)

	if reflect.Array == valueKind.Kind() {

		fetchValueList := fetechValueObject.([]interface{})

		if len(fetchValueList) > 0 {

			dataObject := fetchValueList[0]

			if dataObject != nil {

				fetchValue = reflect.ValueOf(dataObject).String()

			}

		}

	} else if fetechValueObject != nil {

		fetchValue = reflect.ValueOf(fetechValueObject).String()

	}

	fetchFuncInfo := compute.FetchFuncInfo(featureInfo.Condition)

	switch fetchFuncInfo.FuncName {

	case "DictModelEnum":
		dictValue = dict.FetchDictInfo(fetchFuncInfo.Args, fetchValue)

	case "DictOther":
		dictValue = dict.FetchOtherInfo(fetchFuncInfo.FuncName, fetchFuncInfo.Args, fetchValue)

	case "DictChanenlInfo1":
		dictValue = dict.FetchDictInfo(fetchFuncInfo.FuncName, level1channelId(fetchValue))

	case "DictChanenlInfo2":

		dictValue = dict.FetchDictInfo(fetchFuncInfo.FuncName, level2channelId(fetchValue))

	case "DictCarInfo":

		dictValue = fetchMakeId(fetchValue)

	case "DictSereisLevelId":

		dictValue = fetchSereisLevelId(fetchValue)

	default:
		dictValue = fetchValue
	}

	if tools.IsEmpty(dictValue) {

		dictValue = fetchValue

	}

	return dictValue

}

func level1channelId(chnlCode string) string {

	dimChannInfo := sql.DimChannelInfo{ChnlCode: chnlCode}

	dimChannInfo.FetchEntity()

	return dimChannInfo.ChnlLevel1Name

}

func level2channelId(chnlCode string) string {

	dimChannInfo := sql.DimChannelInfo{ChnlCode: chnlCode}

	dimChannInfo.FetchEntity()

	return dimChannInfo.ChnlLevel2Name

}

func fetchMakeId(styleId string) string {

	dimCarStyleExt := sql.DimCarStyleExt{StyleId: styleId}

	dimCarStyleExt.FetchEntity()

	return dimCarStyleExt.MakeId

}

func fetchSereisLevelId(styleId string) string {

	dimCarStyleExt := sql.DimCarStyleExt{StyleId: styleId}

	dimCarStyleExt.FetchEntity()

	return dimCarStyleExt.SeriesLevelId

}
