package provider

import (
	"fmt"
	"strconv"
	"strings"
)

func (provider FormulaProvier) TwoNumPercent() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	modelInfoMap := provider.ModelInfoMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	calvalue := defaultValue

	if tools.IsNotEmpty(feature.Name) {

		featureNames := tools.SplitComma(feature.Name)

		featureValues := provider.CalRelFeatureValues(featureNames, modelInfoMap, dataMap)

		value1 := featureValues[0]

		value2 := featureValues[1]

		if tools.IsNotEmpty(value1) {

			value1 = strings.Replace(value1, "-999", "", 1)

		}

		if tools.IsNotEmpty(value2) {

			value2 = strings.Replace(value2, "-999", "", 1)

		}

		decimal1, _ := strconv.ParseFloat(value1, 64)

		decimal2, _ := strconv.ParseFloat(value2, 64)

		if decimal2 > 0 {

			calvalue = fmt.Sprintf("%.6f", decimal1/decimal2)

		}

	}

	return calvalue
}
