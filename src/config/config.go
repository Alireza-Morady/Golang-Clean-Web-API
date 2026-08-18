package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Password PasswordConfig
	Cors     CorsConfig
	Logger   LoggerConfig
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
}

type CorsConfig struct {
	AllowOrigins string
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 string
	DialTimeout        time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
}

type PasswordConfig struct {
	IncludeChars     string
	IncludeDigits    string
	MinLength        int
	MaxLength        int
	IncludeUppercase string
	IncludeLowercase string
}

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatal("Errorin LoadConfig %v", err)
	}
	cfg, err := ParserConfig(v)
	if err != nil {
		log.Fatal("Errorin ParseConfig %v", err)
	}
	return cfg
}

func ParserConfig(v *viper.Viper) (*Config, error) {
    var cfg Config

    log.Printf("Redis from Viper: %#v", v.Get("redis"))
    log.Printf("Redis host: %v", v.Get("redis.host"))
    log.Printf("Redis port: %v", v.Get("redis.port"))
    log.Printf("Redis password: %v", v.Get("redis.password"))

    err := v.Unmarshal(&cfg)
    if err != nil {
        log.Printf("Unable to parse config: %v", err)
        return nil, err
    }

    log.Printf("Redis after unmarshal: %+v", cfg.Redis)

    return &cfg, nil
}

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "config/config-docker"
	} else if env == "production" {
		return "/config/config-production"
	} else {
		return "../config/config-development"
	}
}
