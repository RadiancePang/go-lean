package provider

import (
	"strconv"
)

func (provider FormulaProvier) LimitCount() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	calvalue := defaultValue

	modelInfoMap := provider.ModelInfoMap

	multiFeatureInfo := provider.CalMultiFeature(feature.Name, modelInfoMap)

	reportTime := provider.PraseCalendar(multiFeatureInfo.CalendarFeature, modelInfoMap, dataMap)

	conditionList := PraseCondition(feature.Condition)

	fetchMapList := provider.FilterConditionList(multiFeatureInfo.FeatureName, reportTime, conditionList, dataMap)

	if fetchMapList != nil {

		if tools.IsNotEmpty(feature.DefaultCondition) {

			defaultConditionList := PraseCondition(feature.DefaultCondition)

			defaultMapList := provider.FilterConditionList(multiFeatureInfo.FeatureName, reportTime, defaultConditionList, dataMap)

			if len(defaultMapList) > 0 {

				calvalue = strconv.Itoa(len(fetchMapList))

			}

		} else {

			if tools.IsNotEmpty(feature.FilterField) {

				filterList := provider.FetchDataList(feature.FilterField, fetchMapList)

				calvalue = strconv.Itoa(len(filterList))

			} else {

				calvalue = strconv.Itoa(len(fetchMapList))

			}

		}

	}

	return calvalue

}
