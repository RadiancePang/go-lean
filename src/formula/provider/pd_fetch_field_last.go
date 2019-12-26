package provider

import (
	"time"
)

func (provider FormulaProvier) FetchFieldLast() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	maxValue := defaultValue

	conditionList := PraseCondition(feature.Condition)

	fetchMapList := provider.FilterConditionList(feature.Name, time.Time{}, conditionList, dataMap)

	if fetchMapList != nil {

		filterList := provider.FetchFieldDataList(feature.FilterField, feature.FetchField, fetchMapList)

		fetchFieldInfo := compute.CalSortNumFetch(defaultValue, filterList, true)

		if tools.IsNotEmpty(fetchFieldInfo.FetchValue) {

			maxValue = fetchFieldInfo.FetchValue

		}

	}

	return maxValue

}
