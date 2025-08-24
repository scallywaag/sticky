package env

const (
	StickyEnvVar = "STICKY_ENV"
)

type StickyEnv string

const (
	EnvProd StickyEnv = "prod"
	EnvDev  StickyEnv = "dev"
	EnvTest StickyEnv = "test"
)
