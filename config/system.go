package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	Network       string `mapstructure:"network" json:"network" yaml:"network"`
	Address       string `mapstructure:"address" json:"address" yaml:"address"`
	TlsKey        string `mapstructure:"tls-key" json:"tlsKey" yaml:"tls-key"`
	TlsPem        string `mapstructure:"tls-pem" json:"tlsPem" yaml:"tls-pem"`
	Director      string `mapstructure:"director" json:"director" yaml:"director"`
}
