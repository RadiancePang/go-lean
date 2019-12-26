package provider

import (
	"go-learn/components/tools"
	"go-learn/formula/compute"
)

func (provider FormulaProvier) MultiAddNum() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	modeInfoMap := provider.ModelInfoMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	diff := defaultValue

	if tools.IsNotEmpty(feature.Name) {

		featureValueList := provider.CalMultiFeatureStr(feature.Name, modeInfoMap, dataMap)

		diff = compute.CalTotal(defaultValue, feature.Type, featureValueList)

	}

	return diff

}
