package provider

import (
	"go-learn/components/tools"
	"go-learn/formula/compute"
)

func (provider FormulaProvier) OverduePeriodMax() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	maxValue := defaultValue

	fetchDataListob := tools.IterMap(feature.Name, dataMap)

	if fetchDataListob != nil {

		fetchDataList := fetchDataListob.([]interface{})

		filterList := provider.CalMaxOverDuePeriod(feature.Condition, fetchDataList)

		maxValue = compute.CalSortNum(defaultValue, filterList, true)

	}

	return maxValue

}
