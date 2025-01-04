package config

import (
	"LABSAM-WEB3-BACKEND/internal/constants"
	"LABSAM-WEB3-BACKEND/internal/logging"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	// ConfigurationFilePath is the path to the configuration file.
	ConfigurationFilePath = "config/app_config.yml"
)

// JwtConfig represents the jwt tokens configuration.
type JwtConfig struct {
	Secret                 string `mapstructure:"secret"`
	AccessTokenLifetimeMs  int    `mapstructure:"access_token_lifetime_ms"`
	RefreshTokenLifetimeMs int    `mapstructure:"refresh_token_lifetime_ms"`
	IdTokenLifetimeMs      int    `mapstructure:"id_token_lifetime_ms"`
}

// DatabaseConfig represents the database configuration.
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	SSL      bool   `mapstructure:"ssl"`
}

// RateLimit represents the rate limit configuration for the middleware.
type RateLimit struct {
	RequestPerSecond int `mapstructure:"request_per_second"`
	BurstSize        int `mapstructure:"burst_size"`
}

// AppConfig represents the application configuration structure.
type AppConfig struct {
	ServerPort                  string         `mapstructure:"server_port"`
	Database                    DatabaseConfig `mapstructure:"database"`
	JWT                         JwtConfig      `mapstructure:"jwt"`
	RateLimit                   RateLimit      `mapstructure:"rate_limit"`
	LogLevel                    string         `mapstructure:"log_level"`
	AllowedOrigins              []string       `mapstructure:"allowed_origins"`
	AllowAllOrigins             bool           `mapstructure:"allow_all_origins"`
	Environment                 string         `mapstructure:"environment"`
	MaxLoginAttempts            int            `mapstructure:"max_login_attempts"`
	LoginAttemptLockoutPeriodMs int            `mapstructure:"login_attempt_lockout_period_ms"`
}

var (
	cfg     *AppConfig
	cfgOnce sync.Once
)

// initConfig initializes the Viper instance and reads the configuration file.
func initConfig(l logging.Logger) {
	viper.SetConfigFile(ConfigurationFilePath)
	viper.SetConfigType("yaml")

	// Read configuration file
	if err := viper.ReadInConfig(); err != nil {
		l.LogError("Error reading config file", err)
		os.Exit(1)
	}

	// Watch for changes in the configuration file and reload automatically
	viper.WatchConfig()
}

// GetConfig retrieves the application configuration, loading it if necessary.
func GetConfig(l logging.Logger) *AppConfig {
	cfgOnce.Do(func() {
		initConfig(l)

		if err := viper.Unmarshal(&cfg); err != nil {
			l.LogError("Unable to decode into struct", err)
			os.Exit(1)
		}

		setLogLevel(cfg.LogLevel, l)
		validateConfig(cfg, l)
		l.LogInfo("Configuration loaded successfully")
	})

	return cfg
}

// setLogLevel sets the logging level based on the provided log level string.
func setLogLevel(logLevel string, l logging.Logger) {
	switch logLevel {
	case constants.LogLevelInfo:
		l.SetLevel(logging.LoggerLevel(logrus.InfoLevel))
	case constants.LogLevelWarn:
		l.SetLevel(logging.LoggerLevel(logrus.WarnLevel))
	case constants.LogLevelError:
		l.SetLevel(logging.LoggerLevel(logrus.ErrorLevel))
	case constants.LogLevelDebug:
		l.SetLevel(logging.LoggerLevel(logrus.DebugLevel))
	default:
		l.SetLevel(logging.LoggerLevel(logrus.ErrorLevel))
	}
}

// validateConfig validates the given AppConfig configuration.
func validateConfig(cfg *AppConfig, l logging.Logger) {
	var validationErrors []string

	// Validate the server port
	validationErrors = append(validationErrors, validateServerPort(cfg.ServerPort)...)

	// Validate the database configuration
	validationErrors = append(validationErrors, validateDatabaseConfig(cfg.Database)...)

	// Validate the JWT configuration
	validationErrors = append(validationErrors, validateJWTConfig(cfg.JWT)...)

	// Validate allowed origins
	validationErrors = append(validationErrors, validateAllowedOrigins(cfg.AllowedOrigins, l)...)

	// Validate environment
	validationErrors = append(validationErrors, validateEnvironment(cfg.Environment)...)

	// Validate max login attempts
	validationErrors = append(validationErrors, validateMaxLoginAttempts(cfg.MaxLoginAttempts)...)

	// Validate login attempt lockout period
	validationErrors = append(validationErrors, validateLoginAttemptLockout(cfg.LoginAttemptLockoutPeriodMs)...)

	// Check if there are any validation errors
	if len(validationErrors) > 0 {
		for _, err := range validationErrors {
			l.LogError(err)
		}
		os.Exit(1)
	}
}

func validateServerPort(serverPort string) []string {
	if serverPort == "" {
		return []string{"the server port is not properly configured, it should not be empty"}
	}
	return nil
}

func validateDatabaseConfig(db DatabaseConfig) []string {
	var errors []string
	if db.Host == "" || db.Port == "" || db.User == "" || db.Password == "" {
		errors = append(errors, "the master database configuration is not properly set, fields should not be empty")
	}
	return errors
}

func validateJWTConfig(jwt JwtConfig) []string {
	var errors []string
	if jwt.Secret == "" {
		errors = append(errors, "the jwt secret is not properly set, it should not be empty")
	}
	if jwt.AccessTokenLifetimeMs <= 0 || jwt.RefreshTokenLifetimeMs <= 0 || jwt.IdTokenLifetimeMs <= 0 {
		errors = append(errors, "the jwt token lifetimes are not properly set, they should be greater than 0")
	}
	return errors
}

func validateAllowedOrigins(allowedOrigins []string, l logging.Logger) []string {
	var errors []string
	if allowedOrigins == nil || len(allowedOrigins) == 0 {
		errors = append(errors, "the allowed origins is not properly set, it should not be empty")
	}

	for _, origin := range allowedOrigins {
		if ok, err := isValidOrigin(origin); !ok || err != nil {
			errors = append(errors, fmt.Sprintf("the allowed origin '%s' is not valid %s", origin, err))
		}
	}

	if len(errors) == 0 && len(allowedOrigins) > 0 {
		l.LogWarn("No allowed origins provided. This might cause issues.")
	}

	return errors
}

func validateEnvironment(environment string) []string {
	var errors []string
	if environment == "" {
		errors = append(errors, "the environment is not properly configured, it should not be empty")
	}
	if environment != constants.EnvironmentProduction && environment != constants.EnvironmentDevelopment {
		errors = append(errors, fmt.Sprintf("the environment is not properly configured, it should be either '%s' or '%s'", constants.EnvironmentProduction, constants.EnvironmentDevelopment))
	}
	return errors
}

func validateMaxLoginAttempts(maxLoginAttempts int) []string {
	if maxLoginAttempts <= 0 {
		return []string{"the max login attempts is not properly configured, it should be greater than 0"}
	}
	return nil
}

func validateLoginAttemptLockout(lockoutPeriodMs int) []string {
	if lockoutPeriodMs <= 0 {
		return []string{"the login attempt lockout period is not properly configured, it should be greater than 0"}
	}
	return nil
}

// isValidOrigin checks if a string is a valid domain or IP address.
func isValidOrigin(origin string) (bool, error) {
	if !strings.HasPrefix(origin, "http://") && !strings.HasPrefix(origin, "https://") {
		return false, errors.New("origin should start with http:// or https://")
	}
	_, err := url.ParseRequestURI(origin)
	return err == nil, err
}
