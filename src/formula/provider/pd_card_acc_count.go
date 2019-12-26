package provider

import (
	"strconv"
)

func (provider FormulaProvier) CardAccCount() string {

	feature := provider.FeatureInfo

	modelInfoMap := provider.ModelInfoMap

	dataMap := provider.DataMap

	countStr := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	fetchKeys := tools.SplitComma(feature.Name)

	conditionList := PraseCondition(feature.Condition)

	reportTimeFeature, _ := modelInfoMap.Load(fetchKeys[0])

	if reportTimeFeature != nil {

		reportCalendar := provider.PraseCalendar(reportTimeFeature.(face.Feature), modelInfoMap, dataMap)

		fetchDataList := provider.FilterConditionList(fetchKeys[1], reportCalendar, conditionList, dataMap)
		if fetchDataList != nil {

			countStr = strconv.Itoa(len(fetchDataList))

		}

	}

	return countStr

}
