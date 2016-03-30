package models

import (
    "github.com/lexteam/kabosu/modules"
)

func AutoMigrate() {
    modules.DB.AutoMigrate(
        &Service{},
        &Build{},
        &Artifact{},
    )
}
