package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File
	RunMode string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	DbCfg DbConfig
)

type DbConfig struct {
	Type string
	User string
	Pass string
	Host string
	Name string
	TablePrefix string
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
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadDb() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}
	DbCfg.Type = sec.Key("TYPE").MustString("sqlite3")
	DbCfg.User = sec.Key("USER").MustString("root")
	DbCfg.Pass = sec.Key("PASS").MustString("root")
	DbCfg.Name = sec.Key("NAME").MustString("blog")
	DbCfg.TablePrefix = sec.Key("TABLE_PREFIX").MustString("blog_")
}