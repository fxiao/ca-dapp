package conf

import (
	"log"
	"time"

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
