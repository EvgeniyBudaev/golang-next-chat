package config

import (
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/logger"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Port                string `envconfig:"PORT"`
	LoggerLevel         string `envconfig:"LOGGER_LEVEL"`
	Host                string `envconfig:"HOST"`
	DBPort              string `envconfig:"DB_PORT"`
	DBUser              string `envconfig:"DB_USER"`
	DBPassword          string `envconfig:"DB_PASSWORD"`
	DBName              string `envconfig:"DB_NAME"`
	DBSSlMode           string `envconfig:"DB_SSLMODE"`
	JWTSecret           string `envconfig:"JWT_SECRET"`
	JWTIssuer           string `envconfig:"JWT_ISSUER"`
	JWTAudience         string `envconfig:"JWT_AUDIENCE"`
	CookieDomain        string `envconfig:"COOKIE_DOMAIN"`
	Domain              string `envconfig:"DOMAIN"`
	BaseUrl             string `envconfig:"KEYCLOAK_BASE_URL"`
	Realm               string `envconfig:"KEYCLOAK_REALM"`
	ClientId            string `envconfig:"KEYCLOAK_CLIENT_ID"`
	ClientSecret        string `envconfig:"KEYCLOAK_CLIENT_SECRET"`
	RealmRS256PublicKey string `envconfig:"KEYCLOAK_REALM_RS256_PUBLIC_KEY"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := godotenv.Load(); err != nil {
		logger.Log.Debug("error in method godotenv.Load", zap.Error(err))
		return nil, err
	}
	err := envconfig.Process("GNCHAT", &cfg)
	if err != nil {
		logger.Log.Debug("error in method envconfig.Process", zap.Error(err))
		return nil, err
	}
	return &cfg, nil
}
