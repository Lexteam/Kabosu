package build

import (
    "os"
    "log"
    "bytes"
    "os/exec"
    "path/filepath"
    "github.com/lexteam/kabosu/models"
    "github.com/lexteam/kabosu/modules"
    "github.com/lexteam/kabosu/utils"
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

            var build models.Build = models.Build{
                Log: string(cmdOutput.Bytes()),
                Service: service,
            }
            modules.DB.Create(&build)

            artifacts, err := config.GetSection("artifacts")
            if err != nil {
                if artifacts.HasKey("download") {
                    utils.CopyFile(service.Directory + string(filepath.Separator) + artifacts.Key("download").String(),
                        getStorageDir(buildID) + string(filepath.Separator) + artifacts.Key("download").String())

                    modules.DB.Create(&models.Artifact{
                        Name: artifacts.Key("download").String(),
                        Location: getStorageDir(buildID) +
                            string(filepath.Separator) + artifacts.Key("download").String(),
                        Build: build,
                    })
                }
            }
        } else {
            modules.DB.Create(&models.Build{
                Log: string(cmdOutput.Bytes()),
                Service: service,
            })
        }
        return true
    }
    return false
}

func createStorageDir(buildID string) {
    os.Mkdir(getStorageDir(buildID), 0777)
}

func getStorageDir(buildID string) string {
    return modules.CONFIG.Section("storage").Key("DIR").String() + string(filepath.Separator) + buildID
}

func readConfig(dir string) *ini.File {
    config, err := ini.Load([]byte(""), dir + "/kabosu.ini")
    if err != nil {
        log.Fatal("Failed to load config", err)
        return nil
    }

    return config
}
