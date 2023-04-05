package conf

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/hiltpold/lakelandcup-api-gateway/utils"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Configuration struct {
	Host               string `mapstructure:"HOST"`
	Port               string `mapstructure:"PORT"`
	AuthSvcUrl         string `mapstructure:"AUTH_SVC_URL"`
	AccessTokenMaxAge  int    `mapstructure:"ACCESS_TOKEN_MAX_AGE"`
	RefreshTokenMaxAge int    `mapstructure:"REFRESH_TOKEN_MAX_AGE"`
	FranchiseSvcUrl    string `mapstructure:"PROSPECT_SVC_URL"`
}

// Load the environment set with the environment file
func loadEnvironment(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Load(filename)
	} else {
		err = godotenv.Load()
		// handle if .env file does not exist, this is OK
		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}

// LoadGlobal loads configuration from file and environment variables.
func LoadConfig(filename string) (config *Configuration, err error) {
	if err := loadEnvironment(filename); err != nil {
		return nil, err
	}

	fp, fn := filepath.Split(filename)
	fn_splitted := strings.Split(fn, ".")
	configName := strings.Join(fn_splitted[0:len(fn_splitted)-1], ".")
	configType := fn_splitted[len(fn_splitted)-1]

	viper.AddConfigPath(utils.Ternary(fp != "", fp, "."))
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	t := new(Configuration)

	err = viper.Unmarshal(t)

	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)

	if err != nil {
		return nil, err
	}

	return config, nil
}
