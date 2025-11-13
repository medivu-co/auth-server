package envs

var envs serverEnvs

type serverEnvs struct {
	Environment      string `env:"ENVIRONMENT"`
	PostgresDBURL    string `env:"PG_DB_URL"`
	JWTP256PublicKeyHex  string `env:"JWT_P256_PUBLIC_KEY_HEX"`
	JWTP256PrivateKeyHex     string `env:"JWT_P256_PRIVATE_KEY_HEX"`
	JWTExpirationSec int    `env:"JWT_EXP_SEC"`
	RedisAddr        string `env:"REDIS_ADDR"`
	RedisPassword    string `env:"REDIS_PASSWORD"`
	RedisDB          int    `env:"REDIS_DB"`
}

func IsProduction() bool {
	return envs.Environment != "development"
}

func PostgresDBURL() string {
	return envs.PostgresDBURL
}

func JWTP256PublicKeyHex() string {
	return envs.JWTP256PublicKeyHex
}

func JWTP256PrivateKeyHex() string {
	return envs.JWTP256PrivateKeyHex
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
