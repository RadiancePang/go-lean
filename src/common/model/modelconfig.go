package model

import "go-learn/formula/face"

/*特征元数据*/
type FeatureMetaInfo struct {
	DataSetId string `yaml:"dataSetId"`
	FilePath  string `yaml:"filePath"`
}

/*模型元数据*/
type MetaInfo struct {
	Name string `yaml:"name"`

	Version string `yaml:"version"`

	ModelId string `yaml:"modelId"`

	ServiceName string `serviceName`

	DSN string `yaml:"dsn"`

	Features []FeatureMetaInfo `yaml:"features"`

	Status string `yaml:"status"`
}

/*模型配置信息*/
type Config struct {
	Models []MetaInfo `yaml:"models"`
}

/*模型特征map结合*/
type ModelFeature struct {
	Name string

	ModelId string

	Version string

	ServiceName string

	/*模型id+模型特征集合*/
	FeatureMap map[string]face.Feature
}
