package config

import (
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Debug    bool   `yaml:"debug"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Bind9 struct {
		ConfigPath string `yaml:"config_path"`
		ZonesPath  string `yaml:"zones_path"`
		ReloadCmd  string `yaml:"reload_cmd"`
	} `yaml:"bind9"`
	DHCP struct {
		ConfigPath string `yaml:"config_path"`
		LeasesPath string `yaml:"leases_path"`
		RestartCmd string `yaml:"restart_cmd"`
	} `yaml:"dhcp"`
}

func Load() (*Config, error) {
	cfg := &Config{
		Debug: true,
	}

	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnvInt("DB_PORT", 5432)
	cfg.Database.User = getEnv("DB_USER", "admin")
	cfg.Database.Password = getEnv("DB_PASSWORD", "password")
	cfg.Database.Name = getEnv("DB_NAME", "dnsdhcp")

	cfg.Bind9.ConfigPath = getEnv("BIND9_CONFIG_PATH", "/etc/bind/named.conf")
	cfg.Bind9.ZonesPath = getEnv("BIND9_ZONES_PATH", "/etc/bind/zones")
	cfg.Bind9.ReloadCmd = getEnv("BIND9_RELOAD_CMD", "rndc reload")

	cfg.DHCP.ConfigPath = getEnv("DHCP_CONFIG_PATH", "/etc/dhcp/dhcpd.conf")
	cfg.DHCP.LeasesPath = getEnv("DHCP_LEASES_PATH", "/var/lib/dhcp/dhcpd.leases")
	cfg.DHCP.RestartCmd = getEnv("DHCP_RESTART_CMD", "systemctl restart isc-dhcp-server")

	if configFile := os.Getenv("CONFIG_FILE"); configFile != "" {
		data, err := os.ReadFile(configFile)
		if err != nil {
			return nil, err
		}
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
