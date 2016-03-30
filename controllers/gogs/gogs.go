package gogs

import (
    "log"
    "io/ioutil"
    "encoding/json"
    "github.com/google/go-github/github"
    "github.com/lexteam/kabosu/build"
    "gopkg.in/macaron.v1"
)

func GetWebhook(ctx *macaron.Context) {
    body, _ := ioutil.ReadAll(ctx.Req.Body().ReadCloser())

    var res github.PushEvent
    json.Unmarshal(body, &res)

    log.Println("gogs:" + *res.Repo.FullName)
    build.ExecuteBuild("gogs:" + *res.Repo.FullName, *res.HeadCommit.SHA)
}
