package main

import (
    "log"
    "net/http"
    "github.com/lexteam/kabosu/modules"
    "github.com/lexteam/kabosu/models"
    githubController "github.com/lexteam/kabosu/controllers/github"
    gogsController "github.com/lexteam/kabosu/controllers/gogs"
    webController "github.com/lexteam/kabosu/controllers/web"
    "github.com/go-macaron/pongo2"
    "gopkg.in/macaron.v1"
)

func main() {
    // Initialise config
    modules.InitConfig()

    // Initialise database
    modules.InitDatabase()
    models.AutoMigrate()

    // Macaron
    m := macaron.Classic()
    m.Use(pongo2.Pongoer())

    // User interface
    m.Get("/", webController.GetDashboard)
    m.Get("/explore", webController.GetExplore)

    // Webhook
    m.Group("/webhook", func () {
        m.Post("/github", githubController.GetWebhook)
        m.Post("/gogs", gogsController.GetWebhook)
    })

    // Run
    log.Println("Listening on 0.0.0.0:" + modules.CONFIG.Section("web").Key("PORT").String())
    http.ListenAndServe("0.0.0.0:" + modules.CONFIG.Section("web").Key("PORT").String(), m)
}
