package formula

import (
	"go-learn/common/define"
	"go-learn/components/tools"
	"go-learn/formula/face"
	compute "go-learn/formula/task"
	"reflect"
)

/*公式计算*/
func Calcute(defaultCalcutor string, modelInfoMap map[string]face.Feature, paramsMap map[string]interface{}) map[string]string {

	returnDataMap := make(map[string]string)

	if len(modelInfoMap) != 0 {

		computeModel := compute.FeatureComputeModel{}

		computeModel.ModelInfoMap = modelInfoMap

		computeModel.DataMap = paramsMap

		if tools.IsEmpty(defaultCalcutor) {

			defaultCalcutor = define.DefaultFormula

		}

		computeModel.DefaultCalcutor = defaultCalcutor

		computeModel.CallTask()

		computeModel.SyncModelInfoMap.Range(func(key, value interface{}) bool {

			featureInfo := value.(face.Feature)

			if !featureInfo.Hidden {

				returnDataMap[featureInfo.Id] = reflect.ValueOf(featureInfo.Value).String()

			}

			return true

		})

	}

	return returnDataMap

}
