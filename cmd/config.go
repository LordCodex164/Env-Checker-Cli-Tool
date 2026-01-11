package cmd

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() error {
	// Set config file name and paths
	viper.SetConfigName(".envcheck")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")           // Look in current directory
	viper.AddConfigPath("$HOME")       // Look in home directory

	// Set defaults
	viper.SetDefault("schema", "schema.yaml")
	viper.SetDefault("strict", false)
	viper.SetDefault("verbose", false)

	// Read config file (optional - don't fail if not found)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; use defaults and flags
			return nil
		}
		return fmt.Errorf("error reading config file: %w", err)
	}

	return nil
}