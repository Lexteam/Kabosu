package github

import (
    "log"
    "io/ioutil"
    "encoding/json"
    "github.com/google/go-github/github"
    "github.com/lexteam/kabosu/build"
    "gopkg.in/macaron.v1"
)

func GetWebhook(ctx *macaron.Context) {
    // Check it was a push event
    if (ctx.Req.Header.Get("X-GitHub-Event") == "push") {
        body, _ := ioutil.ReadAll(ctx.Req.Body().ReadCloser())

        var res github.PushEvent
        json.Unmarshal(body, &res)

        log.Println("github:" + *res.Repo.FullName)
        build.ExecuteBuild("github:" + *res.Repo.FullName, *res.HeadCommit.SHA)
    }
}
