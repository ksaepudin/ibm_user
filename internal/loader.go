package configs

import (
	"fmt"
	"ibm_users_accsess_management/internal/db"
	"ibm_users_accsess_management/internal/logging"
	"ibm_users_accsess_management/src/shared/util"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Database db.DatabaseList
	Logger   logging.LoggerConfig
}

var configuration Config

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {

	// fmt.Printf("Reading Config %s\n", basepath+"/db")
	viper.AddConfigPath(basepath + "/db")
	viper.SetConfigType("yaml")
	viper.SetConfigName("db.yml")
	err := viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load server config: %v", err))
	}

	// fmt.Printf("Reading Config %s\n", basepath+"/logging")
	viper.AddConfigPath(basepath + "/logging")
	viper.SetConfigName("logger.yml")
	err = viper.MergeInConfig()
	if err != nil {
		log.Println("Failed to load log config: %v", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}

	viper.Unmarshal(&configuration)

	fmt.Println("============================")
	fmt.Println(util.Stringify(configuration))
	fmt.Println("============================")
}

// GetConfig get config
func GetConfig() *Config {
	return &configuration
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(env) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}
