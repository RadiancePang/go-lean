package provider

import (
	"go-learn/components/tools"
	"go-learn/formula/compute"
)

func (provider FormulaProvier) MultiMaxNum() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	modeInfoMap := provider.ModelInfoMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	diff := defaultValue

	if tools.IsNotEmpty(feature.Name) {

		featureValueList := provider.CalMultiFeatureStr(feature.Name, modeInfoMap, dataMap)

		diff = compute.CalSortNum(defaultValue, featureValueList, true)

	}

	return diff

}
