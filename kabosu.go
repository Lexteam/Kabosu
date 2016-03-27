package main

import (
    "log"
    "net/http"
    "github.com/lexteam/kabosu/modules"
    githubController "github.com/lexteam/kabosu/controllers/github"
    "gopkg.in/macaron.v1"
)

func main() {
    // Initialise config
    modules.InitConfig()

    // Macaron
    m := macaron.Classic()

    // Webhook
    m.Post("/github/webhook", githubController.GetWebhook)

    // Run
    log.Println("Listening on 0.0.0.0:" + modules.CONFIG.Section("web").Key("PORT").String())
    http.ListenAndServe("0.0.0.0:" + modules.CONFIG.Section("web").Key("PORT").String(), m)
}
