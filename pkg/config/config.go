package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

const configPath = "./config.yaml"

type Config struct {
	Env         string      `yaml:"env" env-default:"production"`
	HTTPServer  HTTPServer  `yaml:"http_server"`
	GRPCSServer GRPCSServer `yaml:"grpc_server"`
	DB          DB          `yaml:"db"`
	JWT         JWT         `yaml:"jwt"`
	Migrations  Migrations  `yaml:"migrations"`
}

type GRPCSServer struct {
	Network string `yaml:"network" env-default:"tcp"`
	Port    string `yaml:"port" env-default:"50051"`
}

type DB struct {
	Postgres Postgres `yaml:"postgres"`
	Redis    Redis    `yaml:"redis"`
}

type JWT struct {
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl" env:"ACCESS_TOKEN_TTL" env-default:"10s"`
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl" env:"REFRESH_TOKEN_TTL" env-default:"1h"`
	SignKey         string        `env:"SIGN_KEY" env-default:""`
}

type Postgres struct {
	Host     string `yaml:"host" env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
	Username string `yaml:"username" env:"POSTGRES_USER" env-default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"123456"`
	SslMode  string `yaml:"ssl_mode" env:"POSTGRES_SSLMODE" env-default:"disable"`
	DbName   string `yaml:"db_name" env:"POSTGRES_DB" env-default:"web"`
}

type Redis struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"6379"`
	Password string `yaml:"password" env-default:""`
	Db       int    `yaml:"db" env-default:"0"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"0.0.0.0:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Migrations struct {
	Path       string `yaml:"path" env-default:"migrations"`
	Table      string `yaml:"table"`
	SchemaName string `yaml:"schema_name"`
}

func initConfig(cfg *Config) error {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return errors.Wrap(err, fmt.Sprintf("config file %s does not exist", configPath))
	}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return errors.Wrap(err, fmt.Sprintf("reading config file %s failed", configPath))
	}

	return nil
}

func MustLoad() *Config {
	cfg := &Config{}

	err := initConfig(cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
