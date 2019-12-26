package provider

import (
	"reflect"
	"strconv"
)

func (provider FormulaProvier) HatchField() string {

	feature := provider.FeatureInfo

	dataMap := provider.DataMap

	modelInfoMap := provider.ModelInfoMap

	defaultValue := provider.CreditDefaultStr(feature.Type, feature.DataSetId)

	var hatchFieldValue string

	fetchFuncInfo := compute.FetchFuncInfo(feature.Condition)

	var args []string

	if len(fetchFuncInfo.Args) > 0 {

		args = tools.SplitComma(fetchFuncInfo.Args)

		feildValueList := provider.CalRelFeatureValues(args, modelInfoMap, dataMap)

		hatchFieldValue = praseHatchFeild(fetchFuncInfo.FuncName, feildValueList)

	} else {

		hatchFeature, ok := modelInfoMap.Load(feature.Name)

		if ok {

			hatchFeatureIds := []string{hatchFeature.(face.Feature).Id}
			hatchFeatureValues := provider.CalRelFeatureValues(hatchFeatureIds, modelInfoMap, dataMap)

			hatchFieldValue = praseHatchFeild(fetchFuncInfo.FuncName, hatchFeatureValues)

		}

	}

	if !tools.IsNotEmpty(hatchFieldValue) {

		hatchFieldValue = defaultValue

	}

	return hatchFieldValue

}

func praseHatchFeild(funcName string, args []string) string {

	hatchfield := compute.HatchField{FuncName: funcName, Args: args}

	providerRef := reflect.ValueOf(hatchfield)

	methodName := tools.Camel(hatchfield.FuncName)

	providerMethod := providerRef.MethodByName(methodName)

	results := providerMethod.Call(nil)

	fetchResult := results[0]

	var retvalue string

	if fetchResult.IsValid() {

		retvalue = strconv.FormatInt(fetchResult.Int(), 10)

	}

	return retvalue

}
