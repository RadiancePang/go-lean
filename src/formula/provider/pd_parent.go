package provider

import (
	"reflect"
	"sync"
)

/*公共公式提供者*/
type FormulaProvier struct {
	FeatureInfo face.Feature

	ModelInfoMap *sync.Map

	DataMap map[string]interface{}
}

func (provider FormulaProvier) Formula() string {

	defer components.CatchPanic()

	providerRef := reflect.ValueOf(provider)

	methodName := tools.Camel(provider.FeatureInfo.Formula)

	providerMethod := providerRef.MethodByName(methodName)

	results := providerMethod.Call(nil)

	return results[0].String()

}

func (provider FormulaProvier) CreditDefaultStr(metaType string, modelId string) string {

	var defalutValue string

	if modelId == define.ModelYxScore {

		defalutValue = define.FetchInitValue(define.ModelYxScore)

	} else {

		defalutValue = define.FetchInitValue(metaType)

	}

	return defalutValue

}

func PraseCondition(condition string) []face.Condition {

	conditions := tools.SplitSemicolon(condition)

	var fetchCondtions []face.Condition

	for _, fetchCon := range conditions {

		if tools.IsNotEmpty(fetchCon) {

			operChar, fetchCons := tools.SplitEqual(fetchCon)

			if len(fetchCons) > 1 {

				operCondition := face.Condition{}

				operCondition.Key = fetchCons[0]

				operCondition.Value = fetchCons[1]

				operCondition.Operator = operChar

				fetchCondtions = append(fetchCondtions, operCondition)

			}

		}

	}

	return fetchCondtions

}

func (provider FormulaProvier) CalRelationFeature(modelInfoMap *sync.Map, paramsMap map[string]interface{}, features ...face.Feature) []face.Feature {

	var returnFeautres []face.Feature

	for _, feature := range features {

		if tools.IsNotEmpty(feature.Id) {

			calFlag := false

			fetchValue := feature.Value

			fetchValueKind := reflect.ValueOf(fetchValue)

			if fetchValue == nil {

				calFlag = true

			} else if fetchValueKind.Kind() == reflect.String {

				calFlag = true

			}

			if calFlag {

				formulaProvider := FormulaProvier{

					ModelInfoMap: modelInfoMap,

					DataMap: paramsMap,

					FeatureInfo: feature,
				}

				result := formulaProvider.Formula()

				feature.Value = result

				modelInfoMap.Store(feature.Id, feature)

			}

			returnFeautres = append(returnFeautres, feature)

		}

	}

	return returnFeautres

}

func (provider FormulaProvier) CalRelFeatureValues(feildIds []string, modelInfoMap *sync.Map, paramsMap map[string]interface{}) []string {

	var featureList []face.Feature

	var featureValueList []string

	for _, feildId := range feildIds {

		if fetchFeildInfo, ok := modelInfoMap.Load(feildId); ok && fetchFeildInfo != nil {

			featureList = append(featureList, fetchFeildInfo.(face.Feature))

		} else if compute.HasDigit(feildId) {

			featureValueList = append(featureValueList, feildId)

		}

	}

	featureList = provider.CalRelationFeature(modelInfoMap, paramsMap, featureList...)

	for _, feature := range featureList {

		if feature.Value != nil {

			featureValueList = append(featureValueList, feature.Value.(string))

		}

	}

	return featureValueList

}
