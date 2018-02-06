package conf

import (
	"strings"

	"github.com/spf13/viper"
)

var Cfg Configs

func init() {
	Cfg = NewConfig()
}

type Configs interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	Init()
}

type Config struct {
}

func (v *Config) Init() {
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`yaml`)
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

}

func (v *Config) GetString(key string) string {
	return viper.GetString(key)
}

func (v *Config) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *Config) GetBool(key string) bool {
	return viper.GetBool(key)
}

func NewConfig() Configs {
	v := &Config{}
	v.Init()
	return v
}
