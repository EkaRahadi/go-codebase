package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App         AppConfig
	Filebeat    FilebeatConfig
	HttpServer  HttpServerConfig
	Database    DatabaseConfig
	Redis       RedisConfig
	JWTConfig   JWTConfig
	Otlp        OtlpConfig
	OauthConfig OauthConfig
}

type AppConfig struct {
	AppName     string
	Environment string
	LogLevel    string
}

type FilebeatConfig struct {
	Host string
	Port int
}

type HttpServerConfig struct {
	Host              string
	Port              int
	GracePeriod       int
	MaxUploadFileSize int64
}

type DatabaseConfig struct {
	Postgres              Postgres
	Mysql                 Mysql
	Sslmode               string
	MaxIdleConn           int
	MaxOpenConn           int
	MaxConnLifetimeMinute int
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

type RedisConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type JWTConfig struct {
	AccessSecretKey      string
	RefreshSecretKey     string
	Issuer               string
	AccessTokenLifespan  time.Duration
	RefreshTokenLifespan time.Duration
}

type OtlpConfig struct {
	OtelExporterOtlpMetricsEndpoint string
}

type OauthConfig struct {
	OauthClientId     string
	OauthClientSecret string
}

func parseIntConfig(envKey string) int {
	val, err := strconv.ParseInt(os.Getenv(envKey), 10, 32)
	if err != nil {
		log.Fatal("cannot parse " + envKey)
	}
	return int(val)
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		App:         initAppConfig(),
		Filebeat:    initFilebeatConfig(),
		Database:    initDbConfig(),
		HttpServer:  initHttpServerConfig(),
		Redis:       initRedisConfig(),
		JWTConfig:   initJWTConfig(),
		OauthConfig: initOauthConfig(),
	}
}

func initAppConfig() AppConfig {
	appName := os.Getenv("APP_NAME")
	environment := os.Getenv("APP_ENVIRONMENT")
	logLevel := os.Getenv("APP_LOGLEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	return AppConfig{
		AppName:     appName,
		Environment: environment,
		LogLevel:    logLevel,
	}
}

func initFilebeatConfig() FilebeatConfig {
	host := os.Getenv("FILEBEAT_HOST")
	if host == "" {
		host = "localhost"
	}
	port, err := strconv.ParseInt(os.Getenv("FILEBEAT_PORT"), 10, 32)
	if err != nil {
		port = 7030
	}
	return FilebeatConfig{
		Host: host,
		Port: int(port),
	}
}

func initHttpServerConfig() HttpServerConfig {
	host := os.Getenv("HTTP_SERVER_HOST")

	port, err := strconv.ParseInt(os.Getenv("HTTP_SERVER_PORT"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse HTTP_SERVER_PORT")
	}

	gracePeriod, err := strconv.ParseInt(os.Getenv("HTTP_SERVER_GRACE_PERIOD"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse HTTP_SERVER_GRACE_PERIOD")
	}

	maxUploadFileSizeKb, err := strconv.ParseInt(os.Getenv("HTTP_MAX_UPLOAD_FILE_SIZE_KB"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse HTTP_SERVER_GRACE_PERIOD")
	}

	return HttpServerConfig{
		Host:              host,
		Port:              int(port),
		GracePeriod:       int(gracePeriod),
		MaxUploadFileSize: maxUploadFileSizeKb * 1024,
	}
}

func initDbConfig() DatabaseConfig {
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort, err := strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 10, 32)
	if err != nil {
		postgresPort = 5432
		// log.Fatal("cannot parse POSTGRES_PORT")
	}
	postgresUsername := os.Getenv("POSTGRES_USERNAME")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDbName := os.Getenv("POSTGRES_DBNAME")

	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort, err := strconv.ParseInt(os.Getenv("MYSQL_PORT"), 10, 32)
	if err != nil {
		mysqlPort = 3306
		// log.Fatal("cannot parse MYSQL_PORT")
	}
	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDbName := os.Getenv("MYSQL_DBNAME")

	sslMode := os.Getenv("DB_SSL_MODE")
	maxIdleConn, err := strconv.ParseInt(os.Getenv("DB_MAX_IDLE_CONN"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_MAX_IDLE_CONN")
	}

	maxOpenConn, err := strconv.ParseInt(os.Getenv("DB_MAX_OPEN_CONN"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_MAX_OPEN_CONN")
	}

	connMaxLifetime, err := strconv.ParseInt(os.Getenv("DB_CONN_MAX_LIFETIME"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_CONN_MAX_LIFETIME")
	}

	return DatabaseConfig{
		Postgres: Postgres{
			Host:     postgresHost,
			Port:     int(postgresPort),
			Username: postgresUsername,
			Password: postgresPassword,
			DbName:   postgresDbName,
		},
		Mysql: Mysql{
			Host:     mysqlHost,
			Port:     int(mysqlPort),
			Username: mysqlUsername,
			Password: mysqlPassword,
			DbName:   mysqlDbName,
		},
		Sslmode:               sslMode,
		MaxIdleConn:           int(maxIdleConn),
		MaxOpenConn:           int(maxOpenConn),
		MaxConnLifetimeMinute: int(connMaxLifetime),
	}
}

func initRedisConfig() RedisConfig {
	port, err := strconv.ParseInt(os.Getenv("REDIS_PORT"), 10, 32)
	if err != nil {
		port = 6379
		// log.Fatal("cannot parse REDIS_PORT")
	}

	return RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     int(port),
		Password: os.Getenv("REDIS_PASSWORD"),
		Username: os.Getenv("REDIS_USERNAME"),
	}
}

func initJWTConfig() JWTConfig {
	accessSecretKey := os.Getenv("ACCESS_SECRET_KEY")
	refreshSecretKey := os.Getenv("REFRESH_SECRET_KEY")
	issuer := os.Getenv("JWT_ISSUER")
	accessLifespan, err := strconv.ParseInt(os.Getenv("JWT_ACCESS_TOKEN_LIFESPAN"), 10, 32)
	if err != nil {
		accessLifespan = 5
	}
	refreshLifespan, err := strconv.ParseInt(os.Getenv("JWT_REFRESH_TOKEN_LIFESPAN"), 10, 32)
	if err != nil {
		refreshLifespan = 24
	}
	return JWTConfig{
		AccessSecretKey:      accessSecretKey,
		RefreshSecretKey:     refreshSecretKey,
		Issuer:               issuer,
		AccessTokenLifespan:  time.Minute * time.Duration(accessLifespan),
		RefreshTokenLifespan: time.Hour * time.Duration(refreshLifespan),
	}
}

func initOtlpConfig() OtlpConfig {
	OtelExporterOtlpMetricsEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_METRICS_ENDPOINT")
	if OtelExporterOtlpMetricsEndpoint == "" {
		OtelExporterOtlpMetricsEndpoint = "localhost:4317"
	}
	return OtlpConfig{
		OtelExporterOtlpMetricsEndpoint: OtelExporterOtlpMetricsEndpoint,
	}
}

func initOauthConfig() OauthConfig {
	oauthClientId := os.Getenv("OAUTH_CLIENT_ID")
	oauthClientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	return OauthConfig{
		OauthClientId:     oauthClientId,
		OauthClientSecret: oauthClientSecret,
	}
}
