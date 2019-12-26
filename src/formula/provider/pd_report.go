package provider

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

func (provider FormulaProvier) DefaultFilterConditionList(name string, reportCreateTime time.Time, conditions []face.Condition,
	dataMap map[string]interface{}, defaultConditions []face.Condition) []map[string]interface{} {

	if defaultConditions != nil && len(defaultConditions) > 0 {

		defaultMapList := provider.FilterConditionList(name, reportCreateTime, defaultConditions, dataMap)

		if len(defaultMapList) == 0 {

			return nil

		}

	}

	return provider.FilterConditionList(name, reportCreateTime, conditions, dataMap)
}

func (provider FormulaProvier) FilterConditionList(name string, reportCreateTime time.Time, conditions []face.Condition,
	dataMap map[string]interface{}) []map[string]interface{} {

	fetchMapListOb := tools.IterMap(name, dataMap)

	if fetchMapListOb == nil {

		return nil

	}

	fetchMapList := fetchMapListOb.([]interface{})

	var fetchMapobList []map[string]interface{}

	if len(fetchMapList) == 0 {

		return nil

	}

	var returnList []map[string]interface{}

	computeFlag := false

	if conditions != nil && len(conditions) > 0 {

		computeFlag = true

	} else {

		computeFlag = false

	}

	for _, fetchMapob := range fetchMapList {

		var choiceList []bool

		fetchMap := fetchMapob.(map[string]interface{})

		fetchMapobList = append(fetchMapobList, fetchMap)

		if computeFlag {

			for _, condition := range conditions {

				flag := false

				targetValue := tools.IterMap(condition.Key, fetchMap)

				var sourceStates []string

				if tools.IsNotEmpty(condition.Value) {

					sourceStates = tools.SplitComma(condition.Value)

				}

				isEqual := false

				if tools.IsNotOper(condition.Operator) {

					isEqual = true

				}

				if targetValue != nil {

					targetValueStr := targetValue.(string)

					if tools.IsFunc(condition.Value) {

						funcFlag := compute.ProcessFunc(condition.Value, reportCreateTime, targetValue)

						if funcFlag != nil {

							flag = funcFlag.(bool)

						}

					} else {

						flag = true

						for _, sourceState := range sourceStates {

							if (targetValueStr == sourceState) == isEqual {

								flag = false

								continue

							}

						}

					}

				}

				if flag {

					choiceList = append(choiceList, flag)

				}

			}

			if len(choiceList) == len(conditions) {

				returnList = append(returnList, fetchMap)
			}

		}

	}

	if !computeFlag {

		returnList = fetchMapobList

	}

	return returnList

}

func (provider FormulaProvier) FetchDataList(key string, dataList []map[string]interface{}) []string {

	var fetchDataList []string

	for _, fetchap := range dataList {

		calValue := tools.IterMap(key, fetchap)

		if calValue != nil {

			calValueKind := reflect.ValueOf(calValue)

			if reflect.Array == calValueKind.Kind() || reflect.Slice == calValueKind.Kind() {

				fetchObjectList := calValue.([]interface{})

				for _, fetchObject := range fetchObjectList {

					if fetchObject != nil {

						fetchObjectStr := fetchObject.(string)

						fetchDataList = append(fetchDataList, tools.DealMoney(fetchObjectStr))

					}

				}

			} else {

				fetchDataList = append(fetchDataList, tools.DealMoney(calValue.(string)))

			}

		}

	}

	return fetchDataList

}

func (provider FormulaProvier) FetchFieldDataList(key string, fetchKey string, dataList []map[string]interface{}) []face.FetchFieldInfo {

	var fetchDataList []face.FetchFieldInfo

	for _, fetchap := range dataList {

		fetchFieldInfo := face.FetchFieldInfo{}

		calValue := tools.IterMap(key, fetchap)

		fetchFieldInfo.FetchValue = tools.DefaultStr(calValue, define.MetaTypeString)

		if tools.IsNotEmpty(fetchKey) {

			fetchValue := tools.IterMap(fetchKey, fetchap)

			fetchFieldInfo.FetchValue = tools.DefaultStr(fetchValue, define.MetaTypeString)

		}

		fetchDataList = append(fetchDataList, fetchFieldInfo)

	}

	return fetchDataList

}

func (provider FormulaProvier) CalMultiFeature(name string, modelInfoMap *sync.Map) face.MultiFeatureInfo {

	fetchKeys := tools.SplitComma(name)

	var multiFeatureInfo face.MultiFeatureInfo

	if len(fetchKeys) > 1 {

		reportTimeFeature, ok := modelInfoMap.Load(fetchKeys[0])

		if ok {

			multiFeatureInfo.CalendarFeature = reportTimeFeature.(face.Feature)

			multiFeatureInfo.FeatureName = fetchKeys[1]

		}

	} else {

		multiFeatureInfo.FeatureName = fetchKeys[0]

	}

	return multiFeatureInfo

}

func (provider FormulaProvier) ComputePercentList(name string, compute string, reportCreateTime time.Time, conditions []face.Condition,
	dataMap map[string]interface{}, defalutConditions []face.Condition) []string {

	fetchMapList := provider.DefaultFilterConditionList(name, reportCreateTime, conditions, dataMap, defalutConditions)

	if fetchMapList == nil {

		return nil

	}

	var returnList []string

	for _, fetchMap := range fetchMapList {

		if tools.IsNotEmpty(compute) {

			computeArgs := strings.Split(compute, "/")

			if len(computeArgs) > 1 {

				moleculeStr := tools.DefaultStr(tools.IterMap(computeArgs[0], fetchMap), define.MetaTypeString)

				denominatorStr := tools.DefaultStr(tools.IterMap(computeArgs[1], fetchMap), define.MetaTypeString)

				if tools.IsNotEmpty(moleculeStr) {

					moleculeStr = strings.Replace(moleculeStr, "-999", "", 1)
				}

				if tools.IsNotEmpty(denominatorStr) {

					denominatorStr = strings.Replace(denominatorStr, "-999", "", 1)

				}

				molecule, _ := strconv.ParseFloat(moleculeStr, 64)

				divisor, _ := strconv.ParseFloat(denominatorStr, 64)

				if divisor > molecule {

					percent := fmt.Sprintf("%.6f", molecule/divisor)

					returnList = append(returnList, percent)

				}

			}

		}

	}

	return returnList
}
