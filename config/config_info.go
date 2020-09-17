package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config is configuration for Server
type Config struct {
	// Meaningful values are recommended (eg. production, development, staging, release/123, etc)
	Environment string

	// Turns on some debug functionality
	Debug bool

	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string

	//Version tells the configuration vs application version
	Version string

	//StoragePath , where to store these documents
	StoragePath string
}

// Configure configures some defaults in the Viper instance.
func configure(v *viper.Viper, p *pflag.FlagSet) {
	v.AllowEmptyEnv(true)
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath("./")
	pflag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", "my-app")
		pflag.PrintDefaults()
	}
	_ = v.BindPFlags(p)

	v.Set("serviceName", "my-app")

	// Global configuration
	v.SetDefault("environment", "production")
	v.SetDefault("debug", false)
	v.SetDefault("shutdownTimeout", 15*time.Second)

	// Server configuration
	p.String("grpcport", ":8090", "App HTTP server address")
	v.SetDefault("httpport", ":8091")

}

func CreateDir() {
	_, err := os.Stat(Conf.StoragePath)
	if !os.IsNotExist(err) {
		return
	}
	if Conf.StoragePath != "" {
		err := os.Mkdir(Conf.StoragePath, 0755)
		if err != nil {
			log.Println("Failed to create dir ", Conf.StoragePath, " due to err ", err)
		} else {
			log.Println("storage created successfully...")
		}
	}
}
