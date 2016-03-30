package models

import (
    "time"
    "github.com/lexteam/kabosu/modules"
)

type Artifact struct {
    ID int64 `gorm:"primary_key"`
    Name string
    Location string

    Build Build
    BuildID int64

    CreatedAt time.Time
    UpdatedAt time.Time
}

func (a Artifact) GetBuild() Build {
    var build Build
    modules.DB.Model(&a).Related(&build)
    return build
}
