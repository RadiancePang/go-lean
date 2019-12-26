package provider

import (
	"strconv"
)

func (provider FormulaProvier) CaseCloseNum() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	calvalue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	filterListob := tools.IterMap(feature.Name, dataMap)

	if filterListob != nil {

		var dataList []string

		if filterList, ok := filterListob.([]interface{}); ok {

			for _, fetchMapob := range filterList {

				fetchMap := fetchMapob.(map[string]interface{})

				srcDate := tools.IterMap(feature.FetchField, fetchMap)

				srcDateStr := tools.DefaultStr(srcDate, define.MetaTypeString)

				if tools.IsNotEmpty(srcDateStr) {

					_, err := tools.PraseTime(srcDateStr, feature.Formate)

					if err == nil {

						continue

					}

				}

				dataList = append(dataList, srcDateStr)

			}

		}

		calvalue = strconv.Itoa(len(dataList))
	}

	return calvalue
}
