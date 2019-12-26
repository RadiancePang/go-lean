package provider

import (
	"go-learn/components/tools"
	"go-learn/formula/compute"
	"go-learn/formula/face"
)

func (provider FormulaProvier) MathPercentMax() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	modelInfoMap := provider.ModelInfoMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	calValue := defaultValue

	fetchKeys := tools.SplitComma(feature.Name)

	conditionList := PraseCondition(feature.Condition)

	reportTimeFeature, _ := modelInfoMap.Load(fetchKeys[0])

	reportCalendar := provider.PraseCalendar(reportTimeFeature.(face.Feature), modelInfoMap, dataMap)

	defaultConditionList := PraseCondition(feature.DefaultCondition)

	percentList := provider.ComputePercentList(fetchKeys[1],
		feature.Compute, reportCalendar, conditionList, dataMap, defaultConditionList)

	if percentList != nil {

		calValue = compute.CalSortNum(defaultValue, percentList, true)

	}

	return calValue

}
