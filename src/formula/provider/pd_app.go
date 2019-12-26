package provider

import (
	"go-learn/common/define"
	"go-learn/components/tools"
	"go-learn/formula/face"
	"strconv"
)

const (
	keyApplyTime = "apply_time"

	keyCertType = "cert_type"

	keyCertNum = "cert_num"

	certTypeId = "1"
)

func (provider FormulaProvier) CustAge() string {

	var ageStr string

	feildIds := []string{keyApplyTime, keyCertNum, keyCertType}

	provider.CalRelFeatureValues(feildIds, provider.ModelInfoMap, provider.DataMap)

	applyTimeFeauture, ok := provider.ModelInfoMap.Load(keyApplyTime)

	if ok {

		applyTime := applyTimeFeauture.(face.Feature)

		valueStr := tools.DefaultStr(applyTime.Value, define.MetaTypeString)

		if tools.IsNotEmpty(valueStr) {

			appyDate, err := tools.PraseTime(valueStr, "yyyy.MM.dd HH:mm:ss")

			if err == nil {

				currYear := appyDate.Year()

				var certYear int

				var certYearStr string

				certTypeFeature, typeOk := provider.ModelInfoMap.Load(keyCertType)

				certNoFeature, certNoOk := provider.ModelInfoMap.Load(keyCertNum)

				if typeOk && certNoOk {

					certtype := certTypeFeature.(face.Feature).Value.(string)

					certNo := certNoFeature.(face.Feature).Value.(string)

					/*身份证*/
					if certtype == certTypeId && tools.IsNotEmpty(certNo) {

						certYearStr = certNo[6:10]

					}

				}

				if tools.IsNotEmpty(certYearStr) {

					certYear, _ = strconv.Atoi(certYearStr)

				}

				if certYear > 0 {

					diff := currYear - certYear

					ageStr = strconv.Itoa(diff)

				}

			}

		}

	}

	return ageStr

}

func (provider FormulaProvier) ApplyMonth() string {

	var monthStr string

	modelInfoMap := provider.ModelInfoMap

	feildIds := []string{keyApplyTime}

	provider.CalRelFeatureValues(feildIds, modelInfoMap, provider.DataMap)

	applyTime, ok := modelInfoMap.Load(keyApplyTime)

	if ok && applyTime != nil {

		applyTimeOb := applyTime.(face.Feature)

		value := tools.DefaultStr(applyTimeOb.Value, define.MetaTypeString)

		if tools.IsNotEmpty(value) {

			date, err := tools.PraseTime(value, "yyyy.MM.dd hh:mm:ss")

			if err != nil {

				month := int(date.Month())

				monthStr = strconv.Itoa(month)

			}

		}

	}

	return monthStr
}

func (provider FormulaProvier) CustPhone() string {

	var custPhoneStr string

	featureInfo := provider.FeatureInfo

	dataMap := provider.DataMap

	fetchValueObject := tools.IterMap(featureInfo.Name, dataMap)

	if fetchValueObject != nil {

		fetchValue := fetchValueObject.(string)

		if tools.IsNotEmpty(fetchValue) && len(fetchValue) > 3 {

			custPhoneStr = fetchValue[0:3]

		}

	}

	return custPhoneStr

}
