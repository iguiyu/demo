package helpers

const (
	ENV_DEV  = "dev"
	ENV_PROD = "prod"
	ENV_TEST = "test"
	ENV_CI   = "ci"
)

type Env interface {
	IsDev() bool
	IsProd() bool
	IsTest() bool
	IsCI() bool
}

// The base configuration file can be embedded by app's specified configuration
type BaseConfig struct {
	Env         string
	ServiceName string
	DbName      string
	DbDail      string
	RpcPort     string
	HttpPort    string
}

func (this *BaseConfig) IsDev() bool  { return this.Env == ENV_DEV }
func (this *BaseConfig) IsProd() bool { return this.Env == ENV_PROD }
func (this *BaseConfig) IsTest() bool { return this.Env == ENV_TEST }
func (this *BaseConfig) IsCI() bool   { return this.Env == ENV_CI }
