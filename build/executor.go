package build

import (
    "os"
    "log"
    "os/exec"
    "github.com/lexteam/kabosu/models"
    "gopkg.in/ini.v1"
)

func ExecuteBuild(id string) {
    // Check if the repository exists.
    service := models.GetService(id)
    if service.ID != -1 {
        log.Println(service.Directory)

        cmd := exec.Command("git", "pull")
        cmd.Dir = service.Directory
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.Run()

        var config = readConfig(service.Directory)
        if config != nil {
            // has config, now lets run all the build stages
            stages, err := config.GetSection("stages")

            if err != nil {
                // the build stage
                if stages.HasKey("build") {
                    cmd := exec.Command(stages.Key("build").String())
                    cmd.Dir = service.Directory
                    cmd.Stdout = os.Stdout
                    cmd.Stderr = os.Stderr
                    cmd.Run()
                }
            }
        }
    }
}

func readConfig(dir string) *ini.File {
    config, err := ini.Load([]byte(""), dir + "/kabosu.ini")
    if err != nil {
        log.Fatal("Failed to load config", err)
        return nil
    }

    return config
}
