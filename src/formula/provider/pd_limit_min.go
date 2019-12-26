package provider

import (
	"time"
)

func (provider FormulaProvier) LimitMin() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	calvalue := defaultValue

	conditionList := PraseCondition(feature.Condition)

	fetchMapList := provider.FilterConditionList(feature.Name, time.Time{}, conditionList, dataMap)

	if len(fetchMapList) > 0 {

		filterList := provider.FetchDataList(feature.FilterField, fetchMapList)

		calvalue = compute.CalSortNum(defaultValue, filterList, false)

	}

	return calvalue

}
