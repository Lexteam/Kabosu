package modules

import (
    "log"
    "strings"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
    DB *gorm.DB
)

func InitDatabase() {
    var err error

    switch strings.ToLower(CONFIG.Section("database").Key("DB_TYPE").String()) {
    case "sqlite":
        DB, err = gorm.Open("sqlite3", CONFIG.Section("database").Key("PATH").String())
    case "mysql":
        DB, err = gorm.Open("mysql",
            CONFIG.Section("database").Key("USER").String() +
            ":" +
            CONFIG.Section("database").Key("PASS").String() +
            "@tcp(" + CONFIG.Section("database").Key("HOST").String() + ")/" +
            CONFIG.Section("database").Key("NAME").String() +
            "?charset=utf8&parseTime=True")
    }

    if err != nil {
        log.Fatal("Failed to load database!", err)
    }
}
