package provider

import (
	"strconv"
)

func (provider FormulaProvier) CalDiffMonths() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	modelInfoMap := provider.ModelInfoMap

	diff := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	if tools.IsNotEmpty(feature.Name) {

		calendars, ok := provider.CalTwoTimeDiff(feature.Name, modelInfoMap, dataMap)

		if ok {

			actualDiff := tools.DiffMonths(calendars[0], calendars[1], true)

			diff = strconv.FormatInt(actualDiff, 10)

		}

	}

	return diff
}
