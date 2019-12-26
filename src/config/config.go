package config

import (
	"encoding/json"
	"go-learn/common/define"
	"go-learn/common/model"
	"go-learn/formula/face"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var ModelMetaConfig model.Config

var WebMetaConfig model.WebConfig

var ModelFeatureMap map[string]model.ModelFeature

func InitMeta(evnName string) {

	Load("inference/engine/config/yaml/"+evnName+"/web_config.yaml", define.MetaTypeYaml, &WebMetaConfig)

	Load("inference/engine/config/yaml/model_config.yaml", define.MetaTypeYaml, &ModelMetaConfig)

	praseModelFeature()

}

/*yaml文件读取*/
func Load(filename string, metatype string, webconfig interface{}) (err error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	switch metatype {

	case define.MetaTypeJson:
		err = json.Unmarshal(data, webconfig)
	case define.MetaTypeYaml:
		err = yaml.Unmarshal(data, webconfig)

	}

	if err != nil {
		return err
	}

	return nil

}

/*解析模型特征*/
func praseModelFeature() {

	modelFeatureMetaMap := make(map[string]model.ModelFeature)

	for _, modelMetaInfo := range ModelMetaConfig.Models {

		featureMapInfo := make(map[string]face.Feature)

		for _, featureMetaInfo := range modelMetaInfo.Features {

			var modelFeatureInfos []face.Feature

			Load(featureMetaInfo.FilePath, define.MetaTypeJson, &modelFeatureInfos)

			convertFeatureMap(featureMetaInfo.DataSetId, modelFeatureInfos, featureMapInfo)

		}

		modelFeatureOb := model.ModelFeature{

			Name:        modelMetaInfo.Name,
			Version:     modelMetaInfo.Version,
			ModelId:     modelMetaInfo.ModelId,
			ServiceName: modelMetaInfo.ServiceName,
			FeatureMap:  featureMapInfo,
		}

		modelFeatureMetaMap[modelMetaInfo.ModelId] = modelFeatureOb

	}

	ModelFeatureMap = modelFeatureMetaMap

}

/*将特征列表covertmap*/
func convertFeatureMap(dataSetId string, modelFeatureInfos []face.Feature, featureMapInfo map[string]face.Feature) {

	for _, modelfeatureInfo := range modelFeatureInfos {

		modelfeatureInfo.DataSetId = dataSetId

		featureMapInfo[modelfeatureInfo.Id] = modelfeatureInfo

	}
}
