package config

import "github.com/spf13/viper"

type Config struct {
	Port               string `mapstructure:"PORT"`
	AuthSvcUrl         string `mapstructure:"AUTH_SVC_URL"`
	AccessTokenMaxAge  int    `mapstructure:"ACCESS_TOKEN_MAX_AGE"`
	RefreshTokenMaxAge int    `mapstructure:"REFRESH_TOKEN_MAX_AGE"`
	FranchiseSvcUrl    string `mapstructure:"FRANCHISE_SVC_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
