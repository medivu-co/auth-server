package envs

var envs serverEnvs

type serverEnvs struct {
	PostgresDBURL    string `env:"PG_DB_URL"`
	JWTSecretKey     string `env:"JWT_SECRET_KEY"`
	JWTExpirationSec int    `env:"JWT_EXP_SEC"`
	RedisAddr        string `env:"REDIS_ADDR"`
	RedisPassword    string `env:"REDIS_PASSWORD"`
	RedisDB          int    `env:"REDIS_DB"`
}

func PostgresDBURL() string {
	return envs.PostgresDBURL
}

func JWTSecretKey() string {
	return envs.JWTSecretKey
}

func JWTExpirationSec() int {
	return envs.JWTExpirationSec
}

func RedisAddr() string {
	return envs.RedisAddr
}

func RedisPassword() string {
	return envs.RedisPassword
}

func RedisDB() int {
	return envs.RedisDB
}
