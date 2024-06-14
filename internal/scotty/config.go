package scotty

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

// DiscordConfig stores Discord configuration options.  These are typically
// read from a file as part of a Config.
type DiscordConfig struct {
	Token   string `yaml:"token"`
	AppID   string `yaml:"app_id"`
	GuildID string `yaml:"guild_id"`
}

// DatabaseConfig stores database configuration options.
// These are typically read from a file as part of a Config.
type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

// Config stores configuration options for Scotty.  These are
// typically read from a configuration file.
type Config struct {
	Discord  DiscordConfig  `yaml:"discord"`
	Database DatabaseConfig `yaml:"database"`
}

// LoadConfig reads the file at the specified path and unmarshals it
// as a *Config.
func LoadConfig(path string) (*Config, error) {
	config := new(Config)

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %q: %w", path, err)
	}

	err = yaml.Unmarshal(file, &config)

	return config, err
}

// LoadDefaultConfig loads the default configuration file at ~/.config/scotty/config.yaml
func LoadDefaultConfig() (*Config, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}

	configPath := filepath.Join(home, ".config", "scotty", "config.yaml")

	return LoadConfig(configPath)
}
