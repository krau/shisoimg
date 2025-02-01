package dao

import (
	"os"

	"github.com/krau/shisoimg/utils"
	_ "github.com/ncruces/go-sqlite3/embed"
	"github.com/ncruces/go-sqlite3/gormlite"
	"gorm.io/gorm"
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
		PrepareStmt: true,
	})
	if err != nil {
		utils.L.Fatal(err)
		os.Exit(1)
	}
	db.AutoMigrate(&Image{}, &UrlRule{})
}

var urlRules = []UrlRule{}

func Rules() []UrlRule {
	if urlRules != nil {
		return urlRules
	}
	rules, err := GetRules()
	if err == nil {
		urlRules = rules
	}
	return urlRules
}

type Image struct {
	gorm.Model
	Path string
	Md5  string `gorm:"unique"`
}

type UrlRule struct {
	gorm.Model
	Prefix string
	Path   string
}
