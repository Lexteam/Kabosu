package build

import (
    "os"
    "log"
    "os/exec"
    "github.com/lexteam/kabosu/modules"
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

        if _, err := os.Stat(dir + "/kabosu.sh"); err == nil {
            cmd := exec.Command("./kabosu.sh")
            cmd.Dir = dir
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            cmd.Run()
        }
    }
}
