package provider

import (
	"strings"
	"time"
)

func (provider FormulaProvier) OverDueAbove3() string {

	var calvalue = "0"

	feature := provider.FeatureInfo

	conditionList := PraseCondition(feature.Condition)

	dataMap := provider.DataMap

	fetchMapList := provider.FilterConditionList(feature.Name, time.Time{}, conditionList, dataMap)

	if len(fetchMapList) > 0 {

		filterList := provider.FetchDataList(feature.FilterField, fetchMapList)

		for _, fetchValue := range filterList {

			if strings.Contains(fetchValue, "3") {

				calvalue = "1"

				break

			}

		}

	}

	return calvalue

}
