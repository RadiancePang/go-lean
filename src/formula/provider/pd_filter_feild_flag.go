package provider

import (
	"time"
)

func (provider FormulaProvier) FilterFeildFlag() string {

	calvalue := "0"

	featureInfo := provider.FeatureInfo

	dataMap := provider.DataMap

	conditions := PraseCondition(featureInfo.Condition)

	fetchAllMap := provider.FilterConditionList(featureInfo.Name, time.Time{}, conditions, dataMap)

	if len(fetchAllMap) > 0 {

		calvalue = "1"

	}

	return calvalue
}
