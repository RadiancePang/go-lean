package provider

import (
	"go-learn/components/tools"
	"go-learn/formula/compute"
	"strconv"
)

func (provider FormulaProvier) OverduePeriodTotal() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	calValue := defaultValue

	fetchDataListob := tools.IterMap(feature.Name, dataMap)

	if fetchDataListob != nil {

		fetchDataLis := fetchDataListob.([]interface{})

		if len(fetchDataLis) > 0 {

			var diff int

			if tools.IsNotEmpty(feature.Condition) {

				fetchFuncInfo := compute.FetchFuncInfo(feature.Condition)

				diff, _ = strconv.Atoi(fetchFuncInfo.Args)

			}

			filterList := provider.CalOverDuePeriod(diff, fetchDataLis)

			maxValueob := compute.ProcessFunc(feature.Condition, filterList)

			calValue = tools.DefaultStr(maxValueob, feature.Type)

		}

	}

	return calValue

}
