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
            log.Println(modules.CONFIG.Section("services").Key(*res.Repo.FullName).String())

            cmd := exec.Command("git", "pull")
            cmd.Dir = modules.CONFIG.Section("services").Key(*res.Repo.FullName).String()
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            cmd.Run()
        }
    }
}
