package setting

import (
	"log"
	"strings"
	"time"

	"github.com/go-ini/ini"
)

var (
	// Cfg 配置实例
	Cfg *ini.File
	// RunMode 运行模式
	RunMode string
	// HTTPPort 端口号
	HTTPPort int
	// ReadTimeout 读超时
	ReadTimeout time.Duration
	// WriteTimeout 写超时
	WriteTimeout time.Duration
	// DbCfg 数据库配置实例
	DbCfg DbConfig
	// AuthCfg 认证配置实例
	AuthCfg AuthConfig
	// StorageCfg 存储配置
	StorageCfg StorageConfig
)

// DbConfig 数据库配置
type DbConfig struct {
	Type        string
	User        string
	Pass        string
	Host        string
	Name        string
	TablePrefix string
}

// AuthConfig 认证配置
type AuthConfig struct {
	SecretKey string
}

// StorageConfig 存储配置
type StorageConfig struct {
	Endpoint   string
	AccessKey  string
	SecretKey  string
	UseSSL     bool
	BucketName string
}

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadDb()
	LoadAuth()
	LoadStorage()
}

// LoadBase 基础配置
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// LoadServer 服务配置
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// LoadDb 数据库配置
func LoadDb() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}
	DbCfg.Type = sec.Key("TYPE").MustString("mongodb")
	DbCfg.Host = sec.Key("HOST").MustString("mongodb://127.0.0.1:27017")
	DbCfg.User = sec.Key("USER").MustString("root")
	DbCfg.Pass = sec.Key("PASS").MustString("root")
	DbCfg.Name = sec.Key("NAME").MustString("blog")
	DbCfg.TablePrefix = sec.Key("TABLE_PREFIX").MustString("blog_")
}

// LoadAuth 认证配置
func LoadAuth() {
	sec, err := Cfg.GetSection("auth")
	if err != nil {
		log.Fatalf("Fail to get section 'auth': %v", err)
	}
	AuthCfg.SecretKey = sec.Key("SECRET_KEY").MustString("secret_key")
}

// LoadStorage 存储配置
func LoadStorage() {
	sec, err := Cfg.GetSection("storage")
	if err != nil {
		log.Fatalf("Fail to get section 'storage': %v", err)
	}
	endpoints := sec.Key("ENDPOINTS").MustString("Q3AM3UQ867SPQQA43P2F:tfteSlswRu7BJ86wekitnifILbZam1KYY3TG@play.min.io")
	endpointsL := strings.Split(endpoints, "@")
	if len(endpointsL) < 3 {
		log.Fatalf("endpoints config error %v", err)
	} else {
		StorageCfg.Endpoint = endpointsL[0]
		StorageCfg.AccessKey = endpointsL[1]
		StorageCfg.SecretKey = endpointsL[2]
	}
	StorageCfg.UseSSL = sec.Key("USE_SSL").MustBool(false)
	StorageCfg.BucketName = sec.Key("BUCKET_NAME").MustString("tank-blog")
}
