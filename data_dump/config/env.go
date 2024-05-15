package config

type RedisEnv struct {
	Uri      string
	DB       int
	password string
}
type Env struct {
	redis RedisEnv
}

func acquireEnvVars