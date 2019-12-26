package provider

import (
	"time"
)

func (provider FormulaProvier) LimitLast() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	calValue := defaultValue

	conditionList := PraseCondition(feature.Condition)

	fetchMapList := provider.FilterConditionList(feature.Name, time.Time{}, conditionList, dataMap)

	if fetchMapList != nil {

		filterList := provider.FetchDataList(feature.FilterField, fetchMapList)

		if filterList != nil && len(filterList) > 0 {

			calValue = compute.SortDate(filterList, feature.Formate, true)

		}

	}

	return calValue

}
