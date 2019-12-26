package compute

import (
	"math/big"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func HasDigit(source interface{}) bool {

	sourceStr := tools.DefaultStr(source, define.MetaTypeString)

	pattern, _ := regexp.Compile(".*\\d+.*")

	flag := pattern.MatchString(sourceStr)

	return flag
}

func IsOverDue(source interface{}, args []string) bool {

	flag := false

	sourceStr := tools.DefaultStr(source, define.MetaTypeString)

	var argStr string

	var diff int

	if len(args) > 0 {

		argStr = args[0]

		var err error

		diff, err = strconv.Atoi(argStr)

		if err != nil {
			diff = 0
		}

	}

	charArray := strings.Split(sourceStr, "")

	i := len(charArray) - diff

	if i < 0 {

		i = 0
	}

	for j := i; j < len(charArray); j++ {

		fetchCharStr := string(charArray[j])

		if HasDigit(fetchCharStr) {

			flag = true

			break

		}

	}

	return flag

}

func ElementsIn(source interface{}, args []string) bool {

	flag := false

	sourceValue := tools.DefaultStr(source, define.MetaTypeString)

	for _, arg := range args {

		if arg == sourceValue {

			flag = true

			break

		}

	}

	return flag
}

func Greater(source interface{}, args []string) bool {

	flag := false

	sourceValue := tools.DefaultStr(source, define.MetaTypeDecimal)

	targetValue := tools.DefaultStr(args[0], define.MetaTypeDecimal)

	sourceFloat, sourceError := strconv.ParseFloat(sourceValue, 64)

	targetFloat, targetError := strconv.ParseFloat(targetValue, 64)

	if sourceError != nil && targetError != nil {

		if sourceFloat > targetFloat {

			flag = true

		}

	}

	return flag

}

func DiffMonths(sourceObject interface{}, targetObject interface{}, args []string) bool {

	flag := false

	var soureTime time.Time

	soureTime = sourceObject.(time.Time)

	targeStr := tools.DefaultStr(targetObject, define.MetaTypeString)

	if !soureTime.IsZero() && tools.IsNotEmpty(targeStr) {

		targetTime, err := tools.PraseTime(targeStr, args[0])

		if err == nil {

			diff, diffErr := strconv.Atoi(args[1])

			if diffErr == nil {

				targetTime = targetTime.AddDate(0, diff, 0)

				flag = targetTime.After(soureTime)

			}
		}
	}

	return flag
}

//数据排序
//true|最大值 false|最小值
func CalSortNum(defaultValue string, params []string, flag bool) string {

	returnStr := defaultValue

	if len(params) > 0 {

		var numlist []float64

		for _, numStr := range params {

			num, err := strconv.ParseFloat(numStr, 64)

			if err == nil {

				numlist = append(numlist, num)

			}

		}

		if flag {

			sort.Sort(sort.Reverse(sort.Float64Slice(numlist)))

		} else {

			sort.Float64s(numlist)

		}

		if len(numlist) > 0 {

			returnStr = strconv.FormatFloat(numlist[0], 'f', -1, 64)

		}

	}

	return returnStr

}

func CalTotal(defalutValue string, metaType string, parmas []string) string {

	var totalNum float64

	if len(parmas) > 0 {

		bigTotalNum := big.NewFloat(totalNum)

		var paramNum float64

		var err error

		for _, param := range parmas {

			if tools.IsNotEmpty(param) {

				paramNum, err = strconv.ParseFloat(param, 64)

				if err == nil {

					bigParamNum := big.NewFloat(paramNum)

					bigTotalNum.Add(bigParamNum, bigTotalNum)

				}

			}

		}

		totalNum, _ = bigTotalNum.Float64()

	}

	return ConvertMetaType(defalutValue, metaType, totalNum)

}

func ConvertMetaType(defaultValue string, metaType string, metaValue float64) string {

	returnStr := defaultValue
	switch metaType {

	case define.MetaTypeDecimal:
		returnStr = strconv.FormatFloat(metaValue, 'f', -1, 64)
	case define.MetaTypeInt:
		intValue, _ := big.NewFloat(metaValue).Int64()
		returnStr = strconv.FormatInt(intValue, 10)
	default:
		returnStr = strconv.FormatFloat(metaValue, 'f', -1, 64)

	}

	return returnStr
}

func Total(param interface{}) string {

	var totalNum string

	if param != nil {

		dataList := param.([]string)

		totalNum = CalTotal("0", "string", dataList)

	}

	return totalNum

}

func Max(param interface{}) string {

	var maxValue string

	if param != nil {

		dataList := param.([]string)

		maxValue = CalSortNum("0", dataList, true)

	}

	return maxValue

}

func IsZeroCredit(dataMap map[string]interface{}) bool {

	creditKey := "infoSummary.creditCue.creditSummaryCue"

	fetchMap := tools.IterMap(creditKey, dataMap).(map[string]interface{})

	if fetchMap != nil && len(fetchMap) > 0 {

		return true

	} else {

		return false

	}

}
