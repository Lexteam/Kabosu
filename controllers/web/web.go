package web

import (
    "net/http"
    "gopkg.in/macaron.v1"
)

func GetDashboard(ctx *macaron.Context) {
    ctx.HTML(http.StatusOK, "dashboard")
}

func GetExplore(ctx *macaron.Context) {
    ctx.HTML(http.StatusOK, "explore")
}
