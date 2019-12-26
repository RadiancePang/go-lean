package model

/*web context配置信息*/
type WebConfig struct {
	Web struct {
		Context string `yaml:"context"`

		Server struct {
			Port string `yaml:"port"`
		}

		DataSource DataSouceConfig `yaml:"datasource"`
	}

	Rabbitmq RabbitMqConfig `yaml:"rabbitmq"`

	Redis RedisConfig `yaml:"redis"`

	AiModel AiModelConfig `yaml:"aimodel"`

	BrLoanIntentUrl string `yaml:"brLoanIntentUrl"`
}

type DataSouceConfig struct {
	Name string `yaml:"name"`

	Url string `yaml:"url"`

	Active int `yaml:"active"`

	Idle int `yaml:"idle"`

	IdleTimeout int `yaml:"idleTimeout"`
}
