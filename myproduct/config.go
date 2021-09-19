package myproduct

type AppConfig struct {
	AppName  string `env:"APP_NAME" env-default:"MyProduct"`
	AppEnv   string `env:"APP_ENV" env-default:"DEV""`
	Port     string `env:"MY_APP_PORT" env-default:"8081"`
	Host     string `env:"HOST" env-default:"localhost"`
	LogLevel string `env:"LOG_LEVEL" env-default:"ERROR"`
}

var cfg AppConfig
