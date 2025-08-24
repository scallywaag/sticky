package env

const (
	StickyEnvVar = "STICKY_ENV"
	ClrEnvVar    = "CLR"
)

type StickyEnv string
type ClrEnv string

const (
	EnvProd StickyEnv = "prod"
	EnvDev  StickyEnv = "dev"
	EnvTest StickyEnv = "test"

	ClrEnabled  ClrEnv = "true"
	ClrDisabled ClrEnv = "false"
)
