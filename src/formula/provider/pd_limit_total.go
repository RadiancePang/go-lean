package provider

import (
	"time"
)

func (provider FormulaProvier) LimitTotal() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	calvalue := defaultValue

	conditionList := PraseCondition(feature.Condition)

	defaultConditionList := PraseCondition(feature.DefaultCondition)

	fetchMapList := provider.DefaultFilterConditionList(feature.Name,
		time.Time{}, conditionList, dataMap, defaultConditionList)

	if fetchMapList != nil {

		filterList := provider.FetchDataList(feature.FilterField, fetchMapList)

		calvalue = compute.CalTotal(defaultValue, feature.Type, filterList)

	}

	return calvalue

}
