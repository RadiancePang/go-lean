package provider

import (
	"encoding/json"
	"strings"
	"sync"
)

const KeyBrLoanIntent = "key_br_loan_intent"

var mutex sync.Mutex

func (provider FormulaProvier) BrLoanIntent() string {

	var returnValue string

	paramMap := make(map[string]string)

	feature := provider.FeatureInfo

	modelInfoMap := provider.ModelInfoMap

	dataMap := provider.DataMap

	if tools.IsNotEmpty(feature.Condition) {

		conditions := tools.SplitSemicolon(feature.Condition)

		for _, condtion := range conditions {

			fetchFeautreob, _ := modelInfoMap.Load(condtion)

			if fetchFeautreob != nil {

				fetchFeautre := fetchFeautreob.(face.Feature)

				featureValue := fetchFeautre.Value

				featureValueStr := tools.DefaultStr(featureValue, define.MetaTypeString)

				if tools.IsEmpty(featureValueStr) {

					sourceProvider := FormulaProvier{

						ModelInfoMap: modelInfoMap,

						FeatureInfo: fetchFeautre,

						DataMap: dataMap,
					}

					featureValueStr = sourceProvider.Formula()

				}

				paramMap[condtion] = featureValueStr

			} else {

				paramMap[condtion] = ""

			}

		}

	}

	brMdataMap := doRequest(paramMap, dataMap)

	returnValueob := tools.IterMap(feature.Name, brMdataMap)

	returnValue = tools.DefaultStr(returnValueob, define.MetaTypeString)

	return returnValue
}

func doRequest(paramMap map[string]string, dataMap map[string]interface{}) map[string]interface{} {

	mutex.Lock()

	defer mutex.Unlock()

	brMdataMap := make(map[string]interface{})

	if brMdataMapob, ok := dataMap[KeyBrLoanIntent]; ok {

		brMdataMap = brMdataMapob.(map[string]interface{})

	} else if len(paramMap) > 0 {

		url := config.WebMetaConfig.BrLoanIntentUrl

		requestStr := "{\"data\":{\"orderId\":\"apply_no\",\"phone\":\"mobile_phone\",\"certNo\":\"cert_num\",\"name\":\"cust_name\"},\"header\":{\"interId\":\"N0017\",\"tokenKey\":\"risk_01\",\"userId\":\"i_risk01\"}}"

		for paramKey, paramValue := range paramMap {

			paramValue = tools.DefaultStr(paramValue, define.MetaTypeString)

			requestStr = strings.Replace(requestStr, paramKey, paramValue, 1)
		}

		respBytes, resperr := tools.HttpPost([]byte(requestStr), url, 1)

		if resperr == nil {

			jsonerr := json.Unmarshal(respBytes, &brMdataMap)

			if jsonerr == nil {

				if brMdataMap != nil && len(dataMap) > 0 {

					dataMap[KeyBrLoanIntent] = brMdataMap

				}

			}

		}
	}

	return brMdataMap
}
