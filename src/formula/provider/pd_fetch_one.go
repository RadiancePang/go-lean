package provider

import (
	"go-learn/components/tools"
)

type Record struct {
	FilterField string

	TargetField string

	Condition string
}

var initValueMap = map[string]Record{
	"QUERY_RECORD_PERSON": {"queryReason", "sum", "本人查询"},
	"QUERY_LOAN_APPROVE":  {"queryReason", "sum", "贷款审批"},
	"QUERY_CARD_APPROVE":  {"queryReason", "sum", "信用卡审批"},
	"QUERY_LOAN_AFTER":    {"queryReason", "sum", "贷后管理"},
}

/*获取初始化值*/
func FetchInitValue(code string) Record {

	return initValueMap[code]

}

func (provider FormulaProvier) FetchOne() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	defaulValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	returnValue := defaulValue

	fetchDataListob := tools.IterMap(feature.Name, dataMap)

	if fetchDataListob != nil {

		fetchDataList := fetchDataListob.([]interface{})

		if len(fetchDataList) > 0 {

			queryRecord := FetchInitValue(feature.Condition)

			returnValue = filterField(defaulValue, queryRecord, fetchDataList)

		}

	}

	return returnValue
}

func filterField(defaultStr string, queryRecord Record, dataList []interface{}) string {

	returnStr := defaultStr

	for _, dataMapob := range dataList {

		if dataMapob != nil {

			dataMap := dataMapob.(map[string]interface{})

			filterFeild := dataMap[queryRecord.FilterField]

			if filterFeild != nil {

				filterFeildValue := filterFeild.(string)

				if queryRecord.Condition == filterFeildValue {

					targetField := dataMap[queryRecord.TargetField]

					targetFieldValue := tools.DefaultStr(targetField, "")

					if tools.IsNotEmpty(targetFieldValue) {

						returnStr = tools.DealMoney(targetFieldValue)

						break

					}

				}

			}

		}

	}

	return returnStr

}
