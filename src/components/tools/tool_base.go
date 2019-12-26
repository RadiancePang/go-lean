package tools

import (
	"bytes"
	"go-learn/common/define"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	splitCharDot = "."

	splitCharExcl = "!"

	splitCharEqual = "="

	splitCharComma = ","

	splitCharSemicolon = ";"

	funcFlagBegin = "("

	funcFlagEnd = ")"

	arrayFlagBegin = "["

	arrayFlagEnd = "]"
)

/*判断字符串是否为空*/
func IsNotEmpty(param string) bool {

	if len(param) > 0 {

		return true

	} else {

		return false

	}

}

func IsEmpty(param string) bool {

	return !IsNotEmpty(param)

}

/*将字符转换为camel风格*/
func Camel(formula string) string {

	formulaNames := strings.Split(formula, "_")

	var charBuffer bytes.Buffer

	for _, formulaName := range formulaNames {

		formulaName = strings.Title(formulaName)

		charBuffer.WriteString(formulaName)

	}

	return charBuffer.String()

}

func IterMap(name string, dataMap map[string]interface{}) interface{} {

	var mapObject interface{}

	if dataMap != nil && len(dataMap) > 0 {

		iterKeys := SplitDot(name)

		count := 0

		for _, key := range iterKeys {

			mapObjectKind := reflect.ValueOf(mapObject).Kind()

			key = TrimArray(key)

			if mapObject == nil && count == 0 {

				mapObject = dataMap[key]

			} else if (reflect.Array == mapObjectKind || mapObjectKind == reflect.Slice) && mapObject != nil {

				maplist := mapObject.([]interface{})

				var compositeList []interface{}

				for _, mapObject := range maplist {

					mapValue := mapObject.(map[string]interface{})[key]

					if mapValue == nil {

						continue
					}

					mapValueKind := reflect.ValueOf(mapValue).Kind()

					if mapValue != nil && (reflect.Array == mapValueKind || reflect.Slice == mapValueKind) {

						mapValueList := mapValue.([]interface{})

						compositeList = append(compositeList, mapValueList...)

					} else if reflect.String == mapValueKind && mapValue != nil {

						filterValue := mapValue.(string)

						filterValue = DealMoney(filterValue)

						compositeList = append(compositeList, filterValue)

					} else {

						compositeList = append(compositeList, mapValue)

					}
				}

				mapObject = compositeList
			} else if reflect.Map == mapObjectKind {

				fetchMap := mapObject.(map[string]interface{})

				mapObject = fetchMap[key]

				mapObjectKind = reflect.ValueOf(mapObject).Kind()

				if reflect.String == mapObjectKind {
					fetchValue := mapObject.(string)
					fetchValue = DealMoney(fetchValue)
					mapObject = fetchValue
				}
			}

			count++
		}

	}

	return mapObject

}

func SplitDot(param string) []string {

	var paramArray []string

	if IsNotEmpty(param) {

		paramArray = strings.Split(param, splitCharDot)

	}

	return paramArray

}

func SplitComma(param string) []string {

	var paramArray []string

	if IsNotEmpty(param) {

		paramArray = strings.Split(param, splitCharComma)

	}

	return paramArray

}

func SplitSemicolon(param string) []string {

	var paramArray []string

	if IsNotEmpty(param) {

		paramArray = strings.Split(param, splitCharSemicolon)

	}

	return paramArray

}

func DealMoney(param string) string {

	return strings.ReplaceAll(param, splitCharComma, "")

}
func DealStr(param string) string {

	return strings.ReplaceAll(param, "--", "")

}

func IsNotOper(param string) bool {

	if strings.Contains(param, splitCharExcl) {

		return true

	} else {

		return false

	}

}

func SplitEqual(param string) (string, []string) {

	splitChar := splitCharEqual

	if IsNotOper(param) {

		splitChar = splitCharExcl + splitCharEqual

	}

	return splitChar, strings.Split(param, splitChar)

}

func IsFunc(param string) bool {

	funcFlag := false

	if IsNotEmpty(param) {

		if strings.Contains(param, funcFlagBegin) && strings.HasSuffix(param, funcFlagEnd) {

			funcFlag = true

		}

	}

	return funcFlag

}

func SplitFunc(param string) []string {

	funcInfo := strings.SplitN(param, funcFlagBegin, 2)

	funcName := funcInfo[0]

	funcArgs := strings.TrimRight(funcInfo[1], funcFlagEnd)

	funcArray := []string{funcName, funcArgs}

	return funcArray

}

func DefaultStr(param interface{}, metaType string) string {

	defaultValue := define.FetchInitValue(metaType)

	if param != nil {

		var paramstr string

		if paramValueob, ok := param.(string); ok {

			paramstr = paramValueob

		}

		if paramValueob, ok := param.(float64); ok {

			paramstr = strconv.FormatFloat(paramValueob, 'f', -1, 64)

		}

		if IsNotEmpty(paramstr) {

			defaultValue = paramstr

		}

	}

	return defaultValue

}

func TrimArray(param string) string {

	var fieldName string

	if IsNotEmpty(param) {

		if strings.HasPrefix(param, arrayFlagBegin) && strings.HasSuffix(param, arrayFlagEnd) {

			regPattern := regexp.MustCompile("[\\[\\]]")

			fieldName = regPattern.ReplaceAllString(param, "")
		} else {

			fieldName = param

		}

	}

	return fieldName
}

func RedisJoinKey(keys []string) string {

	return strings.Join(keys, define.KeySep)

}

func MergeMap(source map[string]interface{}, target map[string]interface{}) map[string]interface{} {

	if source != nil && target != nil {

		for key, srcValue := range source {

			target[key] = srcValue

		}

		return target

	} else {

		return nil

	}

}
