package provider

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

func (provider FormulaProvier) CalMultiFeatureStr(name string, modeInfoMap *sync.Map, dataMap map[string]interface{}) []string {

	featureNames := tools.SplitComma(name)

	var featureValueList []string

	for _, featureName := range featureNames {

		featureSource, ok := modeInfoMap.Load(featureName)

		if ok && featureSource != nil {

			relationList := provider.CalRelationFeature(modeInfoMap, dataMap, featureSource.(face.Feature))

			featureSourcevalue := relationList[0].Value.(string)

			featureValueList = append(featureValueList, featureSourcevalue)

		}

	}

	return featureValueList

}

func (provider FormulaProvier) CalTwoTimeDiff(name string, modelInfoMap *sync.Map, dataMap map[string]interface{}) ([]time.Time, bool) {

	featureNames := tools.SplitComma(name)

	var calendars []time.Time

	flag := false

	if len(featureNames) == 2 {

		feature1, ok1 := modelInfoMap.Load(featureNames[0])

		feature2, ok2 := modelInfoMap.Load(featureNames[1])

		if ok2 && ok1 {

			sourceCalendar := provider.PraseCalendar(feature1.(face.Feature), modelInfoMap, dataMap)

			targetCalendar := provider.PraseCalendar(feature2.(face.Feature), modelInfoMap, dataMap)

			if !sourceCalendar.IsZero() && !targetCalendar.IsZero() {

				flag = true

				calendars = append(calendars, sourceCalendar, targetCalendar)

			}

		}

	}

	return calendars, flag

}

func (provider FormulaProvier) PraseCalendar(feature face.Feature, modelInfoMap *sync.Map, dataMap map[string]interface{}) time.Time {

	var reportCalendar time.Time

	reportTimeFeatureList := provider.CalRelationFeature(modelInfoMap, dataMap, feature)

	if len(reportTimeFeatureList) > 0 {

		feature = reportTimeFeatureList[0]

		if feature.Value != nil {

			actualValue := feature.Value.(string)

			actualValue = tools.DealStr(actualValue)

			actualValue = strings.Replace(actualValue, "-999", "", 1)

			if tools.IsNotEmpty(actualValue) {

				reportCalendar, _ = tools.PraseTime(actualValue, feature.Formate)

			}

		}
	}

	return reportCalendar
}

func (provider FormulaProvier) CalMaxOverDuePeriod(condition string, dataList []interface{}) []string {

	var returnDataList []string

	args := compute.FetchFuncInfo(condition).Args

	diff, _ := strconv.Atoi(args)

	for _, fetchData := range dataList {

		if fetchData != nil {

			var maxValue string

			fetchDataStr := fetchData.(string)

			var overPeriodList []string

			if tools.IsNotEmpty(fetchDataStr) {

				charArray := strings.Split(fetchDataStr, "")

				var i int

				length := len(charArray)

				if diff <= length {

					i = length - diff

				}

				var count int

				for j := i; j < length; j++ {

					fetchCharStr := charArray[j]

					if compute.HasDigit(fetchCharStr) {

						count++

					} else {

						overPeriodList = append(overPeriodList, strconv.Itoa(count))

						count = 0

					}

				}

				overPeriodList = append(overPeriodList, strconv.Itoa(count))

				maxValue = compute.ProcessFunc(condition, overPeriodList).(string)

			}

			returnDataList = append(returnDataList, maxValue)

		}

	}

	return returnDataList

}

func (provider FormulaProvier) CalOverDuePeriod(diff int, dataList []interface{}) []string {

	var returnDataList []string

	for _, fetchData := range dataList {

		if fetchData != nil {

			var total int

			fetchDataStr := fetchData.(string)

			if tools.IsNotEmpty(fetchDataStr) {

				charArray := strings.Split(fetchDataStr, "")

				var i int

				length := len(charArray)

				if diff <= length {

					i = length - diff

				}

				for j := i; j < length; j++ {

					fetchCharStr := charArray[j]

					if compute.HasDigit(fetchCharStr) {

						total = total + 1

					}

				}

			}

			returnDataList = append(returnDataList, strconv.Itoa(total))

		}

	}

	return returnDataList

}
