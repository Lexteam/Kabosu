package build

import (
    "os"
    "log"
    "os/exec"
    "github.com/lexteam/kabosu/modules"
    "gopkg.in/ini.v1"
)

func ExecuteBuild(id string) {
    // Check if the repository exists.
    if modules.CONFIG.Section("services").HasKey(id) {
        var dir = modules.CONFIG.Section("services").Key(id).String()
        log.Println(dir)

        cmd := exec.Command("git", "pull")
        cmd.Dir = dir
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.Run()

        var config = readConfig(dir)
        if config != nil {
            // has config, now lets run all the build stages
            stages, err := config.GetSection("stages")

            if err != nil {
                // the build stage
                if stages.HasKey("build") {
                    cmd := exec.Command(stages.Key("build").String())
                    cmd.Dir = dir
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
