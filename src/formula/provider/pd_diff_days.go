package provider

import (
	"strconv"
)

func (provider FormulaProvier) CalDaysDiff() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	modelInfoMap := provider.ModelInfoMap

	diff := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	if tools.IsNotEmpty(feature.Name) {

		calendars, ok := provider.CalTwoTimeDiff(feature.Name, modelInfoMap, dataMap)

		if ok {

			actualDiff := tools.DiffDays(calendars[0], calendars[1])

			diff = strconv.FormatInt(actualDiff, 10)

		}

	}

	return diff
}
