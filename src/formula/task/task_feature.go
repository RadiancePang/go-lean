package compute

import (
	"sync"
)

type FeatureComputeModel struct {
	ModelInfoMap map[string]face.Feature

	SyncModelInfoMap sync.Map

	DataMap map[string]interface{}

	DefaultCalcutor string
}

var waitGroup sync.WaitGroup

func (computeModel *FeatureComputeModel) CallTask() {

	for _, featureInfo := range computeModel.ModelInfoMap {

		if tools.IsEmpty(featureInfo.Formula) {

			featureInfo.Formula = computeModel.DefaultCalcutor

		}

		computeModel.SyncModelInfoMap.Store(featureInfo.Id, featureInfo)

	}

	computeModel.SyncModelInfoMap.Range(func(key, value interface{}) bool {

		waitGroup.Add(1)

		go func() {

			defer waitGroup.Done()

			targetFeature := value.(face.Feature)

			formulaProvider := provider.FormulaProvier{

				DataMap: computeModel.DataMap,

				ModelInfoMap: &computeModel.SyncModelInfoMap,
			}

			formulaProvider.FeatureInfo = targetFeature

			result := formulaProvider.Formula()

			targetFeature.Value = result

			computeModel.SyncModelInfoMap.Store(targetFeature.Id, targetFeature)

		}()

		return true

	})

	waitGroup.Wait()

}
