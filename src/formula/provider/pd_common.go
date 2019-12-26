package provider

import (
	"strconv"
)

func (provider FormulaProvier) MapSize() string {

	featureInfo := provider.FeatureInfo

	dataMap := provider.DataMap

	cardNumStr := provider.CreditDefaultStr(featureInfo.Type, featureInfo.DataSetId)

	dataList := tools.IterMap(featureInfo.Name, dataMap)

	if dataList != nil {

		dataArray := dataList.([]interface{})

		cardNumStr = strconv.Itoa(len(dataArray))

	}

	return cardNumStr

}

func (provider FormulaProvier) CalIterDefault() string {

	var value string

	dataMap := provider.DataMap

	featureInfo := provider.FeatureInfo

	if dataMap != nil && len(dataMap) > 0 {

		if tools.IsNotEmpty(featureInfo.Name) {

			paramValue := tools.IterMap(featureInfo.Name, provider.DataMap)

			if paramValue != nil {

				value = tools.DefaultStr(paramValue, featureInfo.Type)

			}

		} else {

			value = featureInfo.DefaultValue

		}

	}

	return value

}
