package build

import (
    "os"
    "log"
    "bytes"
    "os/exec"
    "path/filepath"
    "github.com/lexteam/kabosu/models"
    "github.com/lexteam/kabosu/modules"
    "gopkg.in/ini.v1"
)

func ExecuteBuild(id string, buildID string) bool {
    // Check if the repository exists.
    service := models.GetService(id)
    if service.ID != -1 {
        log.Println(service.Directory)
        // create storage dir (for artifacts)
        createStorageDir(buildID)

        cmdOutput := &bytes.Buffer{}

        cmd := exec.Command("git", "pull")
        cmd.Dir = service.Directory
        cmd.Stdout = cmdOutput
        cmd.Stderr = cmdOutput
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
                    cmd.Stdout = cmdOutput
                    cmd.Stderr = cmdOutput
                    cmd.Run()
                }
            }
        }

        modules.DB.Create(&models.Build{
            Log: string(cmdOutput.Bytes()),
            Service: service,
        })
        return true
    }
    return false
}

func createStorageDir(buildID string) {
    os.Mkdir(modules.CONFIG.Section("storage").Key("DIR").String() + string(filepath.Separator) + buildID, 0777)
}

func readConfig(dir string) *ini.File {
    config, err := ini.Load([]byte(""), dir + "/kabosu.ini")
    if err != nil {
        log.Fatal("Failed to load config", err)
        return nil
    }

    return config
}
