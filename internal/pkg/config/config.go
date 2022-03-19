package config

import (
	"flag"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io"
	"os"
	"sync"
	"time"
)

const (
	defaultLimit           = 3
	defaultPeriod          = time.Duration(3) * time.Minute
	defaultCleanupInterval = time.Duration(30) * time.Second
	defaultWaitPeriod      = time.Duration(180) * time.Second
)

const (
	configFileKey     = "configFile"
	defaultConfigFile = ""
	configFileUsage   = "this is config file path"
)

var (
	once         sync.Once
	cachedConfig *AppConfig
)

type AppConfig struct {
	ServerConfig      ServerConfig      `mapstructure:"app"`
	RateLimiterConfig RateLimiterConfig `mapstructure:"rateLimit"`
}

type ServerConfig struct {
	ServiceName string `mapstructure:"serviceName"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
}

func ProvideAppConfig() (c *AppConfig, err error) {
	once.Do(func() {
		var configFile string
		flag.StringVar(&configFile, configFileKey, defaultConfigFile, configFileUsage)
		flag.Parse()

		var configReader io.ReadCloser
		configReader, err = os.Open(configFile)
		defer configReader.Close() //nolint:staticcheck

		if err != nil {
			return
		}

		c, err = LoadConfig(configReader)
		if err != nil {
			return
		}

		cachedConfig = c
	})

	return cachedConfig, err
}

func (a *AppConfig) GetServerConfig() *ServerConfig {
	return &a.ServerConfig
}

func (a *AppConfig) GetRateLimiterConfig() *RateLimiterConfig {
	return &a.RateLimiterConfig
}

func LoadConfig(reader io.Reader) (*AppConfig, error) {
	var appConfig AppConfig

	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	keysToEnvironmentVariables := map[string]string{
		"app.port":                                    "APP_PORT",
		"rateLimit.rateLimitEnabled":                  "RATE_LIMIT_ENABLED",
		"rateLimit.rateLimitCount":                    "RATE_LIMIT_COUNT",
		"rateLimit.rateLimitPeriodInMinutes":          "RATE_LIMIT_PERIOD_IN_MIN",
		"rateLimit.rateLimitCleanupIntervalInSeconds": "RATE_LIMIT_CLEANUP_INTERVAL_IN_SEC",
		"rateLimit.rateLimitWaitPeriodInSeconds":      "RATE_LIMIT_WAIT_INTERVAL_IN_SEC",
	}

	err := bind(keysToEnvironmentVariables)
	if err != nil {
		return nil, err
	}

	if err := viper.ReadConfig(reader); err != nil {
		return nil, errors.Wrap(err, "Failed to load app config file")
	}

	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, errors.Wrap(err, "Unable to parse app config file")
	}

	return &appConfig, nil
}

func bind(keysToEnvironmentVariables map[string]string) error {
	var bindErrors error

	for key, environmentVariable := range keysToEnvironmentVariables {
		if err := viper.BindEnv(key, environmentVariable); err != nil {
			bindErrors = multierror.Append(bindErrors, err)
		}
	}

	return bindErrors
}
