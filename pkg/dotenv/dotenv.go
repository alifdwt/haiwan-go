package dotenv

import (
	"fmt"

	"github.com/spf13/viper"
)

func Viper(path string) error {
	viper.SetConfigFile(fmt.Sprintf("%s/.env", path))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	return err
}
