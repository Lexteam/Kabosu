package main

import (
    "log"
    "os/exec"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/google/go-github/github"
    "gopkg.in/ini.v1"
    "gopkg.in/macaron.v1"
)

func main() {
    // Config
    config, err := ini.Load([]byte(""), "config.ini")
    if err != nil {
        log.Fatal("Failed to load config", err)
    }

    // Macaron
    m := macaron.Classic()

    // Webhook
    m.Post("/", func (ctx *macaron.Context) {
        // Check it was a push event
        if (ctx.Header().Get("X-GitHub-Event") == "push") {
            body, _ := ioutil.ReadAll(ctx.Req.Body)

            var res github.PushEvent
            json.Unmarshal(body, &res)

            log.Println(*res.Repo.FullName)

            if (config.Section("services").HasKey(*res.Repo.FullName)) {
                cmd := exec.Command("git", "pull")
                cmd.Dir =  config.Section("services").Key(*res.Repo.FullName).String()
                cmd.Start()
            }
        }
    })

    // Run
    log.Println("Listening on 0.0.0.0:" + config.Section("web").Key("PORT").String())
    http.ListenAndServe("0.0.0.0:" + config.Section("web").Key("PORT").String(), m)
}
