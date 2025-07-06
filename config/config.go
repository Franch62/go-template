package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config representa a configuração global do serviço
// Deve ser inicializada via InitConfig()
type Config struct {
	AppName      string
	HTTPPort     int
	DBDSN        string
	Keycloak     KeycloakConfig
	MinIO        MinIOConfig
	Rekor        RekorConfig
	Tracing      TracingConfig
	Prometheus   PrometheusConfig
	LogLevel     string
	TenantHeader string
}

type KeycloakConfig struct {
	URL      string
	Realm    string
	ClientID string
	Secret   string
}

type MinIOConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
	UseSSL    bool
}

type RekorConfig struct {
	URL string
}

type TracingConfig struct {
	OTLPEndpoint string
}

type PrometheusConfig struct {
	Enabled bool
	Path    string
}

var Cfg Config

// InitConfig carrega as configurações usando Viper
func InitConfig() error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Flags
	pflag.String("app_name", "myservice", "Nome do serviço")
	pflag.Int("http_port", 8080, "Porta HTTP")
	pflag.String("db_dsn", "", "DSN do banco de dados")
	pflag.String("log_level", "info", "Nível de log")
	pflag.String("tenant_header", "X-Tenant-ID", "Header de tenant")
	pflag.Parse()
	if err := v.BindPFlags(pflag.CommandLine); err != nil {
		return err
	}

	// .env
	if _, err := os.Stat(".env"); err == nil {
		v.SetConfigFile(".env")
		_ = v.MergeInConfig()
	}

	// config.yaml
	if _, err := os.Stat("config/config.yaml"); err == nil {
		v.SetConfigFile("config/config.yaml")
		_ = v.MergeInConfig()
	} else if _, err := os.Stat("config.yaml"); err == nil {
		v.SetConfigFile("config.yaml")
		_ = v.MergeInConfig()
	}

	if err := v.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("erro ao decodificar config: %w", err)
	}
	return nil
}
