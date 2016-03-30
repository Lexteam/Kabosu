package models

import (
    "time"
    "github.com/lexteam/kabosu/modules"
)

type Build struct {
    ID int64 `gorm:"primary_key"`
    Log string

    Service Service
    ServiceID int64

    Artifacts []Artifact

    CreatedAt time.Time
    UpdatedAt time.Time
}

func (b Build) GetService() Service {
    var service Service
    modules.DB.Model(&b).Related(&service)
    return service
}

func (b Build) GetArtifacts() []Artifact {
    var artifacts []Artifact
    modules.DB.Model(&b).Related(&artifacts)
    return artifacts
}
