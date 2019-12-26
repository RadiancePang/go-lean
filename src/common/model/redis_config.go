package model

type RedisConfig struct {
	Cluster RedisCluster `yaml:"cluster"`

	Pool RedisPool `yaml:"pool"`

	Connection RedisConn `yaml:"connection"`
}

type RedisCluster struct {
	Nodes []string `yaml:"nodes"`
}

type RedisPool struct {
	MaxActive int `yaml:"maxActive"`

	MaxIdle int `yaml:"maxIdle"`

	IdleTimeout int `yaml :"idleTimeout"`

	MaxConnLifetime int `yaml :"maxConnLifetime"`

	TestOnBorrow int `yaml :"testOnBorrow"`
}

type RedisConn struct {
	ReadTimeout int `yaml :"readTimeout"`

	WriteTimeout int `yaml :"writeTimeout"`

	ConnectTimeout int `yaml :"connectTimeout"`
}
