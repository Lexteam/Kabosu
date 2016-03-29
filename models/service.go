package models

import (
    "time"
    "github.com/lexteam/kabosu/modules"
)

type Service struct {
    ID int64 `gorm:"primary_key"`
    Name string
    Directory string

    Builds []Build

    CreatedAt time.Time
    UpdatedAt time.Time
}

func (s Service) GetBuilds() []Build {
    var builds []Build
    modules.DB.Model(&s).Related(&builds)
    return builds
}

func GetService(name string) Service {
    var service Service
    modules.DB.FirstOrInit(&service, "name = ?", name, &Service{ID: -1})
    return service
}
