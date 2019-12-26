package provider

import (
	"strconv"
	"time"
)

func (provider FormulaProvier) MaxId() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	calvalue := defaultValue

	conditionList := PraseCondition(feature.Condition)

	fetchMapList := provider.FilterConditionList(feature.Name, time.Time{}, conditionList, dataMap)

	if len(fetchMapList) > 0 {

		var maxIdList []string

		for _, fetchMap := range fetchMapList {

			dataList := tools.IterMap(feature.FilterField, fetchMap)

			if dataList != nil {

				maxIdList = append(maxIdList, strconv.Itoa(len(dataList.([]interface{}))))

			}

		}

		if len(maxIdList) > 0 {

			calvalue = compute.CalSortNum(defaultValue, maxIdList, true)

		}

	}

	return calvalue

}
