package config

import (
	"errors"
	"flag"
	"strings"

	"github.com/spf13/viper"
	"gitlab.mvalley.com/adam/common/pkg/errs"
)

var DefaultConfigFileName = flag.String("cfn", "config", "name of configs file")
var DefaultConfigFilePath = flag.String("cfp", "./configs/", "path of configs file")

func InitConfiguration(configName string, configPaths []string, config interface{}) error {
	vp := viper.New()
	vp.SetConfigName(configName)
	vp.AutomaticEnv()
	for _, configPath := range configPaths {
		vp.AddConfigPath(configPath)
	}

	if err := vp.ReadInConfig(); err != nil {
		return errs.WithInternalError(err)
	}

	err := vp.Unmarshal(config)
	if err != nil {
		return errs.WithInternalError(err)
	}

	return nil
}

func InitConfig(configFileName, configFilePath *string, c interface{}) error {
	if configFileName == nil {
		return errors.New("config file name is nil")
	}
	if configFilePath == nil {
		return errors.New("config File Path is nil")
	}
	return InitConfiguration(*configFileName, strings.Split(*configFilePath, ","), c)
}
