package github

import (
    "os"
    "log"
    "os/exec"
    "io/ioutil"
    "encoding/json"
    "github.com/lexteam/kabosu/modules"
    "github.com/google/go-github/github"
    "gopkg.in/macaron.v1"
)

func GetWebhook(ctx *macaron.Context) {
    // Check it was a push event
    if (ctx.Req.Header.Get("X-GitHub-Event") == "push") {
        body, _ := ioutil.ReadAll(ctx.Req.Body().ReadCloser())

        var res github.PushEvent
        json.Unmarshal(body, &res)

        log.Println(*res.Repo.FullName)

        if (modules.CONFIG.Section("services").HasKey(*res.Repo.FullName)) {
            var dir = modules.CONFIG.Section("services").Key(*res.Repo.FullName).String()
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
}
