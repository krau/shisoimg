package dao

import (
	"os"
	"time"

	"github.com/krau/shisoimg/utils"
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/ncruces/go-sqlite3/gormlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() {
	if db != nil {
		return
	}
	if err := os.MkdirAll("data", 0755); err != nil {
		utils.L.Fatal(err)
		os.Exit(1)
	}
	var err error
	db, err = gorm.Open(gormlite.Open("data/shisoimg.db"), &gorm.Config{
		Logger:      logger.New(utils.L, logger.Config{Colorful: true, SlowThreshold: time.Second * 5, LogLevel: logger.Error, IgnoreRecordNotFoundError: true, ParameterizedQueries: true}),
		PrepareStmt: true,
	})
	if err != nil {
		utils.L.Fatal(err)
		os.Exit(1)
	}
	db.AutoMigrate(&Image{}, &UrlRule{})
	Rules()
}

var urlRules = []UrlRule{}

func Rules() []UrlRule {
	if len(urlRules) > 0 {
		return urlRules
	}
	rules, err := GetRules()
	if err != nil {
		utils.L.Errorf("Failed to get rules: %v", err)
		return urlRules
	}
	urlRules = rules
	return rules
}

type Image struct {
	gorm.Model
	Path   string
	Md5    string `gorm:"unique"`
	Width  int
	Height int
}

type UrlRule struct {
	gorm.Model
	Prefix string
	Path   string
}
