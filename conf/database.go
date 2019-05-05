package conf

import (
	"log"

	"github.com/go-ini/ini"
)

type Database struct {
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting = &Database{}
var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load(".env.ini")
	if err != nil {
		log.Fatalf("Fail to parse '.env.ini': %v", err)
	}

	d_err := cfg.Section("database").MapTo(DatabaseSetting)
	if d_err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
