package configs

import (
	"log"
	"os"

	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	GRPCServerPort string `mapstructure:"DEPS_GRPC_PORT"`
	GRPCServerAddr string `mapstructure:"DEPS_GRPC_ADDR"`
	MongodbHost    string `mapstructure:"MONGODB_HOST"`

	ApiBaseUrl                 string `mapstructure:"API_BASE_URL"`
	DatabaseName               string `mapstructure:"DATABASE_NAME"`
	HttpClientTimeOutMilliSecs int    `mapstructure:"HTTP_CLIENT_TIMEOUT_MILLISECS"`
	DeputadosIdsCollection     string `mapstructure:"DEPUTADOS_IDS_COLLECTION"`
	DeputadosCollection        string `mapstructure:"DEPUTADOS_COLLECTION"`
	DeputadosPagesUrlBase      string `mapstructure:"DEPUTADO_PAGES_URL"`
}

func NewConfig(base_directory string) (*Config, error) {

	var err error
	if base_directory == "." {
		base_directory, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}
	log.Printf("NewConfig %s", base_directory)

	return loadConfig(base_directory)
}

func loadConfig(path string) (*Config, error) {

	var cfg *Config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(filepath.Join(path, ".env"))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	viper.SetConfigType("json")
	viper.SetConfigFile(filepath.Join(path, "config.json"))
	err = viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	log.Println("Config loaded")
	return cfg, err
}
