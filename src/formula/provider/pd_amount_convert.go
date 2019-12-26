package provider

import (
	"strconv"
)

func (provider FormulaProvier) AmountConvert() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	totalValue := defaultValue

	fetchValueObject := tools.IterMap(feature.Name, dataMap)

	var fetchValue string

	if fetchValueObject != nil {

		fetchValue = tools.DefaultStr(fetchValueObject, feature.Type)

	}

	if tools.IsNotEmpty(feature.Condition) && tools.IsNotEmpty(fetchValue) {

		dividor, err1 := strconv.ParseFloat(feature.Condition, 64)

		molecule, err2 := strconv.ParseFloat(fetchValue, 64)

		if err1 == nil && err2 == nil {

			if dividor > 0 {

				result := molecule / dividor

				totalValue = strconv.FormatFloat(result, 'f', -1, 64)

			} else {

				totalValue = fetchValue

			}

		}

	}

	return totalValue

}
