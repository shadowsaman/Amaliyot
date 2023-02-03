package config

type Config struct {
	HTTPPort string

	PostgresHost           string
	PostgresUser           string
	PostgresDatabase       string
	PostgresPassword       string
	PostgresPort           string
	PostgresMaxConnections int32

	AuthSecretKey string
}

func Load() Config {

	var cfg Config

	cfg.HTTPPort = ":8080"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "samandar"
	cfg.PostgresDatabase = "subscription_service"
	cfg.PostgresPassword = "saman107"
	cfg.PostgresPort = "5432"
	cfg.PostgresMaxConnections = 20

	cfg.AuthSecretKey = "9K+WgNTglA44Hg=="

	return cfg
}
