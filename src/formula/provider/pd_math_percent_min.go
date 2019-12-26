package provider

import (
	"go-learn/components/tools"
	"go-learn/formula/compute"
	"go-learn/formula/face"
)

func (provider FormulaProvier) MathPercentMin() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	modelInfoMap := provider.ModelInfoMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	calValue := defaultValue

	conditionList := PraseCondition(feature.Condition)

	multiFeatureInfo := provider.CalMultiFeature(feature.Name, modelInfoMap)

	reportCalenar := provider.PraseCalendar(multiFeatureInfo.CalendarFeature, modelInfoMap, dataMap)

	var defaultConditionList []face.Condition

	if tools.IsNotEmpty(feature.DefaultCondition) {

		defaultConditionList = PraseCondition(feature.DefaultCondition)

	}

	percentList := provider.ComputePercentList(multiFeatureInfo.FeatureName,
		feature.Compute, reportCalenar, conditionList, dataMap, defaultConditionList)

	calValue = compute.CalSortNum(defaultValue, percentList, false)

	return calValue

}
