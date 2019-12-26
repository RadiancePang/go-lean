package model

type AiModelConfig struct {
	Deploy DeployMeta `yaml:"deploy"`
}

type DeployMeta struct {
	Hosts []string `yaml:"hosts"`
}
